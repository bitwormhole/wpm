package ienv

import (
	"fmt"
	"os"
	"os/user"
	"sort"
	"strings"

	"github.com/bitwormhole/wpm"
	"github.com/starter-go/afs"
	"github.com/starter-go/application"
)

type environment struct {
	facade wpm.Environment
	dd     wpm.DataDir

	baseServerPort    int
	currentServerPort int
	currentUserName   string
	currentUserHome   afs.Path
	currentUserWPMDir afs.Path
	currentExeFile    afs.Path
	filesFolder       afs.Path
	useHTTPS          bool
}

////////////////////////////////////////////////////////////////////////////////

// EnvironmentImpl 。。。
type EnvironmentImpl struct {

	//starter:component
	_as func(wpm.Environment) //starter:as("#")

	FS afs.FS //starter:inject("#")

	ServerPortNumMin int    //starter:inject("${server.port}")
	ServerPortNumMax int    //starter:inject("${server.port.max}")
	ServerProtocol   string //starter:inject("${server.protocol}")

	env *environment
}

func (inst *EnvironmentImpl) _impl() (wpm.Environment, application.Lifecycle) {
	return inst, inst
}

// Life ...
func (inst *EnvironmentImpl) Life() *application.Life {
	return &application.Life{
		OnCreate: inst.init,
	}
}

func (inst *EnvironmentImpl) init() error {
	inst.getEnv()
	return nil
}

func (inst *EnvironmentImpl) getEnv() *environment {

	// get from cache

	env := inst.env
	if env != nil {
		return env
	}

	// load new

	loader := new(envLoader)
	loader.parent = inst
	env = loader.load()
	inst.env = env
	return env
}

// BaseServerPort ...
func (inst *EnvironmentImpl) BaseServerPort() int {
	return inst.getEnv().baseServerPort
}

// CurrentServerPort ...
func (inst *EnvironmentImpl) CurrentServerPort() int {
	return inst.getEnv().currentServerPort
}

// CurrentUserName ...
func (inst *EnvironmentImpl) CurrentUserName() string {
	return inst.getEnv().currentUserName
}

// CurrentUserHome ...
func (inst *EnvironmentImpl) CurrentUserHome() afs.Path {
	return inst.getEnv().currentUserHome
}

// CurrentExecutableFile ...
func (inst *EnvironmentImpl) CurrentExecutableFile() afs.Path {
	return inst.getEnv().currentExeFile
}

// DataDir ...
func (inst *EnvironmentImpl) DataDir() wpm.DataDir {
	return inst.getEnv().dd
}

// UseHTTPS ...
func (inst *EnvironmentImpl) UseHTTPS() bool {
	return inst.getEnv().useHTTPS
}

////////////////////////////////////////////////////////////////////////////////

type dataDirImpl struct {
	env *environment
}

func (inst *dataDirImpl) _impl() wpm.DataDir {
	return inst
}

func (inst *dataDirImpl) Path() afs.Path {
	return inst.env.currentUserWPMDir
}

// 取 ~/.wpm 文件夹下的子目录或文件
func (inst *dataDirImpl) GetPath(path string) afs.Path {
	base := inst.Path()
	return base.GetChild(path)
}

// 取 ~/.wpm/files 文件夹
func (inst *dataDirImpl) Files() afs.Path {
	return inst.env.filesFolder
}

// 取 ~/.wpm/files/:name 文件夹
func (inst *dataDirImpl) GetMediaBucket(name string) afs.Path {
	name = inst.normalizeBucketName(name)
	base := inst.Files()
	return base.GetChild(name)
}

func (inst *dataDirImpl) normalizeBucketName(name string) string {
	b := new(strings.Builder)
	chs := []rune(name)
	for _, ch := range chs {
		if '0' <= ch && ch <= '9' {
			b.WriteRune(ch)
		} else if 'a' <= ch && ch <= 'z' {
			b.WriteRune(ch)
		} else if 'A' <= ch && ch <= 'Z' {
			b.WriteRune(ch)
		} else if ch == '.' || ch == '-' || ch == '_' {
			b.WriteRune(ch)
		}
	}
	if b.Len() == 0 {
		return "media" // default bucket name
	}
	return strings.ToLower(b.String())
}

////////////////////////////////////////////////////////////////////////////////

type envLoader struct {
	parent *EnvironmentImpl
}

func (inst *envLoader) load() *environment {

	e := new(environment)
	dd := new(dataDirImpl)
	steps := make([]func(e *environment) error, 0)

	e.dd = dd
	e.facade = inst.parent
	dd.env = e

	// 注意：这里的加载顺序有讲究 ...
	steps = append(steps, inst.loadAppExeFile)
	steps = append(steps, inst.loadCurrentUserInfo)
	steps = append(steps, inst.loadUserWPMDataDir)
	steps = append(steps, inst.loadFilesDir)
	steps = append(steps, inst.loadServerConfig)

	for _, fn := range steps {
		err := fn(e)
		if err != nil {
			panic(err)
		}
	}

	return e
}

func (inst *envLoader) loadCurrentUserInfo(e *environment) error {
	u, err := user.Current()
	if err != nil {
		return err
	}
	userhome := inst.parent.FS.NewPath(u.HomeDir)
	username := u.Username
	if !userhome.IsDirectory() {
		path := userhome.GetPath()
		return fmt.Errorf("user home dir is not exists, at path [%s]", path)
	}
	e.currentUserName = username
	e.currentUserHome = userhome
	return nil
}

func (inst *envLoader) loadUserWPMDataDir(l *environment) error {
	base := l.currentUserHome
	target := base.GetChild(".wpm")
	l.currentUserWPMDir = target
	return nil
}

func (inst *envLoader) loadFilesDir(l *environment) error {
	base := l.currentUserWPMDir
	target := base.GetChild("files")
	l.filesFolder = target
	return nil
}

func (inst *envLoader) loadAppExeFile(l *environment) error {
	arg0 := ""
	for i, ar := range os.Args {
		if i == 0 {
			arg0 = ar
			break
		}
	}
	if arg0 == "" {
		return fmt.Errorf("exe file path is empty")
	}
	file := inst.parent.FS.NewPath(arg0)
	if !file.IsFile() {
		return fmt.Errorf("no exe file at path [%s]", arg0)
	}
	l.currentExeFile = file
	return nil
}

func (inst *envLoader) loadServerConfig(l *environment) error {

	parent := inst.parent

	// for port ( 根据 user home dir 的位置， 确定 port 的偏移量 )

	portMin := parent.ServerPortNumMin
	portMax := parent.ServerPortNumMax
	if portMin < 1 {
		// check min
		portMin = wpm.DefaultServerPort
	}
	if portMax < portMin {
		// check max
		portMax = portMin
	}

	port := portMin
	homeDir := l.currentUserHome
	homeName := homeDir.GetName()
	namelist := homeDir.GetParent().ListNames()
	sort.Strings(namelist)
	for i, name := range namelist {
		if name == homeName {
			port = portMin + i
			break
		}
	}

	if port > portMax {
		port = portMax
	}

	l.baseServerPort = portMin
	l.currentServerPort = port

	// for protocol
	protocol := parent.ServerProtocol
	l.useHTTPS = (strings.ToLower(protocol) == "https")
	return nil
}

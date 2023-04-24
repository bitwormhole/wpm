package implservice

import (
	"os"
	"strings"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/service"
)

// 定义 app-data 文件名
const (
	AppDataFileDB        = "wpm.db"
	AppDataMainGit       = "main/.git"
	AppDataBackupDumpDir = "backups/dumps"
	AppDataBackupExeDir  = "backups/executables"
)

// AppDataServiceImpl ...
type AppDataServiceImpl struct {
	markup.Component `id:"AppDataService"`

	ProfileService    service.ProfileService    `inject:"#ProfileService"`
	AppRuntimeService service.AppRuntimeService `inject:"#AppRuntimeService"`
	FS                service.FileSystemService `inject:"#FileSystemService"`

	DatabaseName string `inject:"${datasource.wpm.database}"`

	baseAppDataDir afs.Path
}

func (inst *AppDataServiceImpl) _Impl() service.AppDataService {
	return inst
}

func (inst *AppDataServiceImpl) loadAppDataDir() afs.Path {
	profile, err := inst.ProfileService.GetProfile()
	if err != nil {
		panic(err)
	}
	home := inst.FS.Path(profile.Home)
	return home.GetChild(".wpm")
}

func (inst *AppDataServiceImpl) getAppDataDir() afs.Path {
	addir := inst.baseAppDataDir
	if addir == nil {
		addir = inst.loadAppDataDir()
		inst.baseAppDataDir = addir
	}
	return addir
}

func (inst *AppDataServiceImpl) forPath(child string) afs.Path {
	addir := inst.getAppDataDir()
	path := addir.GetChild(child)
	return path
}

func (inst *AppDataServiceImpl) forPathString(child string) string {
	return inst.forPath(child).GetPath()
}

// GetAppDataDirectory 。。。
func (inst *AppDataServiceImpl) GetAppDataDirectory() string {
	return inst.forPathString(".")
}

// GetBackupDumpDirectory 。。。
func (inst *AppDataServiceImpl) GetBackupDumpDirectory() string {
	return inst.forPathString(AppDataBackupDumpDir)
}

// GetSQLiteDBFile ...
func (inst *AppDataServiceImpl) GetSQLiteDBFile() string {
	name := inst.DatabaseName
	am := service.GetAppModule()
	if am != nil {
		version := am.GetVersion()
		sum := inst.getThisExeSha256sum()
		hex := sum.String()[0:7]
		name = strings.ReplaceAll(name, "{{version}}", version)
		name = strings.ReplaceAll(name, "{{hex}}", hex)
	}
	if name == "" {
		name = AppDataFileDB
	}
	return inst.forPathString("databases/" + name)
}

func (inst *AppDataServiceImpl) getThisExeSha256sum() util.Hex {
	hex, _ := inst.AppRuntimeService.GetAppHash()
	str := hex.String()
	if len(str) < 20 {
		hex = "00000000000000000000"
	}
	return hex
}

// GetMainRepositoryPath ...
func (inst *AppDataServiceImpl) GetMainRepositoryPath() string {
	return inst.forPathString(AppDataMainGit)
}

// Ready ...
func (inst *AppDataServiceImpl) Ready() bool {
	dir := inst.forPath(AppDataMainGit)
	return dir.Exists()
}

// Setup ...
func (inst *AppDataServiceImpl) Setup() error {
	dir := inst.forPath(AppDataMainGit)
	vlog.Info("setup main repository at ", dir.GetPath())
	opt := &afs.Options{
		Permission: os.ModePerm,
		Create:     true,
		Mkdirs:     true,
	}
	return dir.Mkdirs(opt)
}

// GetBackupExecutableFile ...
func (inst *AppDataServiceImpl) GetBackupExecutableFile(sum util.Hex) string {
	filename := sum.String()
	filename = strings.TrimSpace(filename)
	if filename == "" {
		filename = "unnamed"
	}
	b := strings.Builder{}
	b.WriteString(AppDataBackupExeDir)
	b.WriteString("/")
	b.WriteString(filename)
	path := b.String()
	return inst.forPathString(path)
}

// GetRoot ...
func (inst *AppDataServiceImpl) GetRoot() afs.Path {
	return inst.forPath("share")
}

// GetPath 根据选项生成对应的路径
func (inst *AppDataServiceImpl) GetPath(opt *service.GetPathOptions) afs.Path {

	tp := opt.Type
	id := opt.ID
	name := opt.Name

	if tp == "" {
		tp = "none"
	}

	if id == "" {
		id = "none/none"
	} else {
		if len(id) < 5 {
			id = "shortids/" + id
		} else {
			const p1size = 2
			p1 := id[0:p1size]
			p2 := id[p1size:]
			id = p1 + "/" + p2
		}
	}

	if name == "" {
		name = "."
	}

	b := strings.Builder{}
	b.WriteString(tp)
	b.WriteString("/")
	b.WriteString(id)
	b.WriteString("/")
	b.WriteString(name)
	return inst.forPath(b.String())
}

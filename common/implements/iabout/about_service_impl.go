package iabout

import (
	"context"
	"crypto/md5"
	"crypto/sha256"

	"runtime"
	"strconv"
	"strings"

	"github.com/bitwormhole/wpm"
	"github.com/bitwormhole/wpm/common/classes/about"
	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/vo"
	"github.com/starter-go/afs"
	"github.com/starter-go/application"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/vlog"
)

// ServiceImpl ...
type ServiceImpl struct {

	//starter:component
	_as func(about.Service) //starter:as("#")

	AC  application.Context //starter:inject("context")
	Env wpm.Environment     //starter:inject("#")
	FS  afs.FS              //starter:inject("#")

	Name      string //starter:inject("${application.name}")
	Title     string //starter:inject("${application.title}")
	Copyright string //starter:inject("${application.copyright}")
	Profile   string //starter:inject("${application.profiles.active}")

	cached *vo.About
}

func (inst *ServiceImpl) _impl() about.Service {
	return inst
}

// GetAboutInfo ...
func (inst *ServiceImpl) GetAboutInfo(ctx context.Context, d *vo.About) error {
	src := inst.cached
	if src == nil {
		src = new(vo.About)
		err := inst.loadInfo(ctx, src)
		if err != nil {
			return err
		}
		inst.cached = src
	}
	*d = *src
	return nil
}

func (inst *ServiceImpl) loadInfo(ctx context.Context, d *vo.About) error {

	username, home := inst.getUserHome()
	sum, exefile := inst.computeAppExeSHA256()

	d.Name = inst.Name
	d.Title = inst.Title
	d.Copyright = inst.Copyright
	d.Profile = inst.Profile
	d.User = username
	d.Home = home
	d.OS = runtime.GOOS
	d.Arch = runtime.GOARCH
	d.WebURL = inst.computeWebServerURL()
	d.Exe = exefile
	d.SHA256SUM = sum

	inst.loadModulesInfo(d)

	mm := &d.MainModule
	d.Module = mm.Name
	d.Version = mm.Version
	d.Revision = mm.Revision

	return nil
}

func (inst *ServiceImpl) getUserHome() (userName string, home string) {
	userName = inst.Env.CurrentUserName()
	home = inst.Env.CurrentUserHome().GetPath()
	return
}

func (inst *ServiceImpl) computeWebServerURL() string {
	return inst.Env.WebServerURL()
}

func (inst *ServiceImpl) computeAppExeSHA256() (retSum lang.Hex, retExeFile string) {

	file := inst.Env.CurrentExecutableFile()
	data, err := file.GetIO().ReadBinary(nil)
	if err != nil {
		vlog.Warn("%s", err.Error())
		return
	}

	sum := sha256.Sum256(data)
	retSum = lang.HexFromBytes(sum[:])
	retExeFile = file.GetPath()
	return
}

func (inst *ServiceImpl) loadModulesInfo(d *vo.About) {
	list1 := inst.AC.GetModules()
	list2 := make([]*dto.Module, 0)
	mm1 := inst.AC.GetMainModule()
	mm2 := inst.convertModule(mm1)
	d.MainModule = *mm2
	for _, m1 := range list1 {
		m2 := inst.convertModule(m1)
		list2 = append(list2, m2)
	}
	d.Modules = list2
}

func (inst *ServiceImpl) convertModule(src application.Module) *dto.Module {
	dst := new(dto.Module)
	dst.Name = src.Name()
	dst.Version = src.Version()
	dst.Revision = src.Revision()

	// make hex name
	hexNameBuilder := strings.Builder{}
	hexNameBuilder.WriteString(dst.Name)
	hexNameBuilder.WriteString("@")
	hexNameBuilder.WriteString(dst.Version)
	hexNameBuilder.WriteString("-r")
	hexNameBuilder.WriteString(strconv.Itoa(dst.Revision))
	bin := []byte(hexNameBuilder.String())
	sum := md5.Sum(bin)
	hex := lang.HexFromBytes(sum[:])

	dst.HexName = hex.String()
	return dst
}

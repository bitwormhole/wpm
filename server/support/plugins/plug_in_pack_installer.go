package plugins

import (
	"fmt"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/components/packs"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
)

const (
	PackTypeZip   = "application/x-wpm-plug-in-zip"
	PackTypeTarGz = "application/x-wpm-plug-in-tar-gz"
	PackTypeOlder = "application/x-wpm-plug-in-package"
)

// WPMPluginInstaller ...
type WPMPluginInstaller struct {
	markup.Component `class:"life packs.InstallerRegistry"`

	HTTPClientExService   service.HTTPClientExService   `inject:"#HTTPClientExService"`
	FileSystemService     service.FileSystemService     `inject:"#FileSystemService"`
	MediaService          service.MediaService          `inject:"#MediaService"`
	ExecutableService     service.ExecutableService     `inject:"#ExecutableService"`
	IntentTemplateService service.IntentTemplateService `inject:"#IntentTemplateService"`

	InstalledFileDAO dao.InstalledFileDAO `inject:"#InstalledFileDAO"`
}

func (inst *WPMPluginInstaller) _Impl() (packs.Installer, packs.InstallerRegistry, application.LifeRegistry) {
	return inst, inst, inst
}

// GetInstallerRegistration ...
func (inst *WPMPluginInstaller) GetInstallerRegistration() *packs.InstallerRegistration {
	return &packs.InstallerRegistration{
		Installer: inst,
	}
}

// GetLifeRegistration ...
func (inst *WPMPluginInstaller) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnInit: inst.init,
	}
}

func (inst *WPMPluginInstaller) init() error {
	return nil
}

// Accept 判断是否支持传入的 IC
func (inst *WPMPluginInstaller) Accept(ic *packs.InstallingContext) bool {
	if ic == nil {
		return false
	}
	ptype := ic.PackType
	return (ptype == PackTypeTarGz || ptype == PackTypeZip || ptype == PackTypeOlder)
}

// Install 安装插件包
func (inst *WPMPluginInstaller) Install(ic *packs.InstallingContext) error {
	task := &myPluginInstallingTask{
		parent: inst,
	}
	return task.Install(ic)
}

// Uninstall 卸载插件包
func (inst *WPMPluginInstaller) Uninstall(ic *packs.InstallingContext) error {
	return fmt.Errorf("no impl")
}

func (inst *WPMPluginInstaller) insertMediae(ic *packs.InstallingContext, list []*dto.Media) error {
	ctx := ic.Context
	ser := inst.MediaService
	opt := &service.MediaOptions{FetchFromSource: true}
	for _, item := range list {
		_, err := ser.Insert(ctx, item, opt)
		if err != nil {
			vlog.Warn(err)
		}
	}
	return nil
}

func (inst *WPMPluginInstaller) insertExecutables(ic *packs.InstallingContext, list []*dto.Executable) error {
	ctx := ic.Context
	ser := inst.ExecutableService
	opt := &service.ExecutableOptions{SkipFileChecking: true}
	for _, item := range list {
		_, err := ser.Insert(ctx, item, opt)
		if err != nil {
			vlog.Warn(err)
		}
	}
	return nil
}

func (inst *WPMPluginInstaller) insertIntentTemplates(ic *packs.InstallingContext, list []*dto.IntentTemplate) error {
	ctx := ic.Context
	ser := inst.IntentTemplateService
	for _, item := range list {
		_, err := ser.Insert(ctx, item)
		if err != nil {
			vlog.Warn(err)
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type myPluginInstallingTask struct {
	parent     *WPMPluginInstaller
	ic         *packs.InstallingContext
	files      []afs.Path
	unpackFunc func(src, dst afs.Path) error
}

func (inst *myPluginInstallingTask) Install(ic *packs.InstallingContext) error {

	inst.ic = ic

	steps := make([]func() error, 0)
	steps = append(steps, inst.prepare)
	steps = append(steps, inst.unzip)
	steps = append(steps, inst.deployFilesToRoot)
	steps = append(steps, inst.importObjects)
	steps = append(steps, inst.scanFiles)
	steps = append(steps, inst.logFiles)

	for _, step := range steps {
		err := step()
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myPluginInstallingTask) prepare() error {
	ptype := inst.ic.PackType
	switch ptype {
	case PackTypeTarGz:
		inst.unpackFunc = utils.UnpackTar
	case PackTypeZip:
		inst.unpackFunc = utils.Unzip
	default:
		return fmt.Errorf("unsupported plug-in package type:%v", ptype)
	}
	return nil
}

func (inst *myPluginInstallingTask) unzip() error {

	// fs := inst.parent .FileSystemService
	// return utils.Unzip(src, dst)

	src := inst.ic.PackFile
	dst := inst.ic.PackDir
	up := inst.unpackFunc
	if up == nil {
		return fmt.Errorf("the un-pack func is nil")
	}
	return up(src, dst)
}

func (inst *myPluginInstallingTask) deployFilesToRoot() error {
	vlog.Warn("todo:", "deployFilesToRoot")
	return nil
}

func (inst *myPluginInstallingTask) importObjects() error {

	const path = "wpm.plug-in/manifest.json"

	ic := inst.ic
	ctx := ic.Context
	httpSer := inst.parent.HTTPClientExService
	manifest := inst.ic.PackDir.GetChild(path)
	url := utils.GetFileURI(manifest)
	opt := &service.HTTPClientOptions{}

	doc, err := httpSer.FetchOnlineDoc(ctx, url, opt)
	if err != nil {
		return err
	}

	inst.parent.insertMediae(ic, doc.Mediae)
	inst.parent.insertExecutables(ic, doc.Executables)
	inst.parent.insertIntentTemplates(ic, doc.IntentTemplates)

	return nil
}

func (inst *myPluginInstallingTask) scanFiles() error {
	path := inst.ic.PackDir
	limit := 99
	return inst.scanPathInto(path, limit)
}

func (inst *myPluginInstallingTask) scanPathInto(p afs.Path, depthLimit int) error {
	if p == nil {
		return nil
	}
	if depthLimit < 0 {
		return fmt.Errorf("the path is too deep: %v", p.GetPath())
	}
	inst.files = append(inst.files, p)
	if !p.IsDirectory() {
		return nil
	}
	children := p.ListChildren()
	for _, ch := range children {
		err := inst.scanPathInto(ch, depthLimit-1)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *myPluginInstallingTask) logFiles() error {
	list := inst.files
	for _, f := range list {
		err := inst.logFile(f)
		if err != nil {
			vlog.Warn(err)
		}
	}
	return nil
}

func (inst *myPluginInstallingTask) logFile(f afs.Path) error {

	if f == nil {
		return nil
	} else if !f.Exists() {
		return nil
	}

	daoIF := inst.parent.InstalledFileDAO
	packURN := inst.ic.Pack.URN
	installation := inst.ic.Installation

	o := &entity.InstalledFile{}
	o.Name = f.GetName()
	o.Path = f.GetPath()
	o.Size = f.GetInfo().Length()
	o.OwnerPackage = packURN
	o.IsDir = f.IsDirectory()
	o.IsFile = f.IsFile()
	o.Installation = installation

	if o.IsFile {
		hex, err := utils.ComputeFileSHA256sumAFS(f)
		if err == nil {
			o.SHA256SUM = hex
		}
	}

	_, err := daoIF.Insert(o)
	return err
}

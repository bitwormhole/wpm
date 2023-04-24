package plugins

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/components/packs"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dbagent"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// PluginServiceImpl ...
type PluginServiceImpl struct {
	markup.Component `id:"SoftwarePackageService"`

	SoftwarePackageDAO dao.SoftwarePackageDAO `inject:"#SoftwarePackageDAO"`

	GormDBAgent dbagent.GormDBAgent `inject:"#GormDBAgent"`

	NamespaceService  service.NamespaceService      `inject:"#NamespaceService"`
	HTTPClient        service.HTTPClientService     `inject:"#HTTPClientService"`
	HTTPClientEx      service.HTTPClientExService   `inject:"#HTTPClientExService"`
	TrashService      service.TrashService          `inject:"#TrashService"`
	IntentTemplateSer service.IntentTemplateService `inject:"#IntentTemplateService"`
	ExecutableSer     service.ExecutableService     `inject:"#ExecutableService"`
	ContentTypeSer    service.ContentTypeService    `inject:"#ContentTypeService"`
	MediaSer          service.MediaService          `inject:"#MediaService"`
	AppDataService    service.AppDataService        `inject:"#AppDataService"`
	// MediaSer1          service.SoftwarePackageService          `inject:"#MediaService"` // self

	InstallerRegistryList []packs.InstallerRegistry `inject:".packs.InstallerRegistry"`

	lockerForGenInstaID sync.Mutex
	lockerForInstall    sync.Mutex
}

func (inst *PluginServiceImpl) _Impl() service.SoftwarePackageService {
	return inst
}

// GetPacksManger ...
func (inst *PluginServiceImpl) GetPacksManger() packs.Manager {
	return inst
}

func (inst *PluginServiceImpl) dto2entity(o1 *dto.SoftwarePackage) (*entity.SoftwarePackage, error) {

	o2 := &entity.SoftwarePackage{}

	o2.ID = o1.ID
	o2.URN = o1.URN
	o2.UUID = o1.UUID
	o2.Referer = o1.Referer
	o2.Installation = o1.Installation

	o2.Name = o1.Name
	o2.ModuleName = o1.ModuleName
	o2.FileName = o1.FileName
	o2.Namespace = o1.Namespace

	o2.Icon = o1.Icon
	o2.Title = o1.Title
	o2.Description = o1.Description

	o2.OS = o1.OS
	o2.Arch = o1.Arch
	o2.Version = o1.Version
	o2.Revision = o1.Revision
	o2.ReleaseAt = o1.ReleaseAt.GetTime()

	o2.ContentType = o1.ContentType
	o2.SHA256SUM = o1.SHA256SUM
	o2.Size = o1.Size

	o2.WebPageURL = o1.WebPageURL
	o2.DownloadURL = o1.DownloadURL
	// o2.ResourceURL = o1.ResourceURL

	return o2, nil
}

func (inst *PluginServiceImpl) entity2dto(o1 *entity.SoftwarePackage) (*dto.SoftwarePackage, error) {

	o2 := &dto.SoftwarePackage{}

	o2.ID = o1.ID
	o2.URN = o1.URN
	o2.UUID = o1.UUID
	o2.Referer = o1.Referer
	o2.Installation = o1.Installation

	o2.Name = o1.Name
	o2.ModuleName = o1.ModuleName
	o2.FileName = o1.FileName
	o2.Namespace = o1.Namespace

	o2.Icon = o1.Icon
	o2.Title = o1.Title
	o2.Description = o1.Description

	o2.OS = o1.OS
	o2.Arch = o1.Arch
	o2.Version = o1.Version
	o2.Revision = o1.Revision
	o2.ReleaseAt = util.NewTime(o1.ReleaseAt)

	o2.ContentType = o1.ContentType
	o2.SHA256SUM = o1.SHA256SUM
	o2.Size = o1.Size

	o2.WebPageURL = o1.WebPageURL
	o2.DownloadURL = o1.DownloadURL
	// o2.ResourceURL = o1.ResourceURL

	o2.Installed = o1.Installation > 0

	return o2, nil
}

func (inst *PluginServiceImpl) entity2dtoList(src []*entity.SoftwarePackage) ([]*dto.SoftwarePackage, error) {
	dst := make([]*dto.SoftwarePackage, 0)
	for _, item := range src {
		item2, err := inst.entity2dto(item)
		if err != nil {
			return nil, err
		}
		dst = append(dst, item2)
	}
	return dst, nil
}

func (inst *PluginServiceImpl) prepareSave(o *entity.SoftwarePackage) error {

	ss := &utils.Strings{}
	urnb := strings.Builder{}

	name := ss.TrimToLower(o.Name)
	ns := ss.TrimToLower(o.Namespace)
	ver := ss.TrimToLower(o.Version)
	rev := "r" + strconv.Itoa(o.Revision)
	os := ss.TrimToLower(o.OS)
	arch := ss.TrimToLower(o.Arch)

	urnb.WriteString(ns)
	urnb.WriteString("#")
	urnb.WriteString(name)
	urnb.WriteString("-")
	urnb.WriteString(ver)
	urnb.WriteString("-")
	urnb.WriteString(os)
	urnb.WriteString("-")
	urnb.WriteString(arch)
	urnb.WriteString("-")
	urnb.WriteString(rev)

	o.ModuleName = ns + "#" + name
	o.URN = dxo.NewSoftwarePackageURN(urnb.String())

	return nil
}

// Find ...
func (inst *PluginServiceImpl) Find(ctx context.Context, id dxo.SoftwarePackageID) (*dto.SoftwarePackage, error) {
	dao := inst.SoftwarePackageDAO
	o1, err := dao.Find(id)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o1)
}

// ListAll ...
func (inst *PluginServiceImpl) ListAll(ctx context.Context) ([]*dto.SoftwarePackage, error) {
	dao := inst.SoftwarePackageDAO
	list1, err := dao.ListAll()
	if err != nil {
		return nil, err
	}
	return inst.entity2dtoList(list1)
}

// ListByModuleName ...
func (inst *PluginServiceImpl) ListByModuleName(ctx context.Context, moduleName string) ([]*dto.SoftwarePackage, error) {
	dao := inst.SoftwarePackageDAO
	srclist, err := dao.ListByModuleName(moduleName)
	if err != nil {
		return nil, err
	}
	return inst.entity2dtoList(srclist)
}

// Insert ...
func (inst *PluginServiceImpl) Insert(ctx context.Context, o1 *dto.SoftwarePackage) (*dto.SoftwarePackage, error) {

	o2, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}

	inst.prepareSave(o2)

	dao := inst.SoftwarePackageDAO
	o3, err := dao.Insert(o2)
	if err != nil {
		return nil, err
	}

	return inst.entity2dto(o3)
}

// UpdateItem ...
func (inst *PluginServiceImpl) UpdateItem(ctx context.Context, id dxo.SoftwarePackageID, o1 *dto.SoftwarePackage) (*dto.SoftwarePackage, error) {

	o2, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}

	inst.prepareSave(o2)

	dao := inst.SoftwarePackageDAO
	o3, err := dao.Update(id, o2)
	if err != nil {
		return nil, err
	}

	return inst.entity2dto(o3)
}

// UpdateList ...
func (inst *PluginServiceImpl) UpdateList(ctx context.Context) error {

	// 下载更新的包信息
	up := &myPakcageListUpdater{
		packageSer:    inst,
		namespaceSer:  inst.NamespaceService,
		httpclientSer: inst.HTTPClientEx,
	}
	err := up.Fetch(ctx)
	if err != nil {
		return err
	}

	// 清除过时的包信息
	cl := &myPakcageListCleaner{
		packageSer:    inst,
		namespaceSer:  inst.NamespaceService,
		httpclientSer: inst.HTTPClientEx,
	}
	err = cl.Clean(ctx)
	if err != nil {
		return err
	}

	// 清除数据库中已经软删除的对象
	err = inst.TrashService.Clean()
	if err != nil {
		return err
	}

	return up.Save(ctx)
}

// Remove ...
func (inst *PluginServiceImpl) Remove(ctx context.Context, id dxo.SoftwarePackageID) error {
	dao := inst.SoftwarePackageDAO
	return dao.Remove(id)
}

func (inst *PluginServiceImpl) generateInstallationID() dxo.InstallationID {
	inst.lockerForGenInstaID.Lock()
	defer func() {
		inst.lockerForGenInstaID.Unlock()
	}()
	time.Sleep(time.Second)
	n := util.Now().Int64()
	return dxo.InstallationID(n)
}

// Install ...
func (inst *PluginServiceImpl) Install(ctx context.Context, id dxo.SoftwarePackageID) error {

	// dao := inst.SoftwarePackageDAO
	// fetch
	p, err := inst.Find(ctx, id)
	if err != nil {
		return err
	}
	vlog.Warn("install package:", p.ID, " ", p.URN)

	installer := myPakcageInstaller{
		context: ctx,

		SoftwarePackageSer: inst,
		HTTPClient:         inst.HTTPClient,
		HTTPClientEx:       inst.HTTPClientEx,
		IntentTemplateSer:  inst.IntentTemplateSer,
		ExecutableSer:      inst.ExecutableSer,
		ContentTypeSer:     inst.ContentTypeSer,
		MediaSer:           inst.MediaSer,
		AppDataSer:         inst.AppDataService,
	}
	installer.installation = inst.generateInstallationID()
	p, err = installer.Install(p)
	if err != nil {
		return err
	}

	// update
	p.Installed = true
	p.Installation = installer.installation

	// save
	_, err = inst.UpdateItem(ctx, id, p)
	if err != nil {
		return err
	}

	return nil
}

// Uninstall ...
func (inst *PluginServiceImpl) Uninstall(ctx context.Context, id dxo.SoftwarePackageID) error {

	p, err := inst.Find(ctx, id)
	if err != nil {
		return err
	}

	vlog.Warn("uninstall package:", p.ID, " ", p.URN)

	uni := myPakcageUninstaller{
		context:      ctx,
		agent:        inst.GormDBAgent,
		installation: p.Installation,
	}
	err = uni.Uninstall()
	if err != nil {
		return err
	}

	p.Installed = false
	p.Installation = 0
	_, err = inst.UpdateItem(ctx, id, p)
	return err
}

// GetInstaller ...
func (inst *PluginServiceImpl) GetInstaller(ic *packs.InstallingContext) (packs.Installer, error) {
	if ic == nil {
		return nil, fmt.Errorf("param:ic(InstallingContext) is nil")
	}
	src := inst.InstallerRegistryList
	for _, r1 := range src {
		r2 := r1.GetInstallerRegistration()
		installer := r2.Installer
		if installer != nil {
			if installer.Accept(ic) {
				return installer, nil
			}
		}
	}
	return nil, fmt.Errorf("no installer for " + ic.String())
}

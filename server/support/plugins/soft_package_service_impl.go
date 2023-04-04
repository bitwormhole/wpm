package plugins

import (
	"context"
	"strconv"
	"strings"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/dao"
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

	NamespaceService  service.NamespaceService  `inject:"#NamespaceService"`
	HTTPClientService service.HTTPClientService `inject:"#HTTPClientService"`
}

func (inst *PluginServiceImpl) _Impl() service.SoftwarePackageService {
	return inst
}

func (inst *PluginServiceImpl) dto2entity(o1 *dto.SoftwarePackage) (*entity.SoftwarePackage, error) {

	o2 := &entity.SoftwarePackage{}

	o2.ID = o1.ID
	o2.URN = o1.URN
	o2.UUID = o1.UUID

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
	o2.ResourceURL = o1.ResourceURL

	o2.Installed = o1.Installed

	return o2, nil
}

func (inst *PluginServiceImpl) entity2dto(o1 *entity.SoftwarePackage) (*dto.SoftwarePackage, error) {

	o2 := &dto.SoftwarePackage{}

	o2.ID = o1.ID
	o2.URN = o1.URN
	o2.UUID = o1.UUID

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
	o2.ResourceURL = o1.ResourceURL

	o2.Installed = o1.Installed

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
	up := &myPakcageListUpdater{
		packageSer:    inst,
		namespaceSer:  inst.NamespaceService,
		httpclientSer: inst.HTTPClientService,
	}
	return up.Update(ctx)
}

// Remove ...
func (inst *PluginServiceImpl) Remove(ctx context.Context, id dxo.SoftwarePackageID) error {
	dao := inst.SoftwarePackageDAO
	return dao.Remove(id)
}

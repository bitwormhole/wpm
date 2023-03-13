package mediae

import (
	"context"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// MediaServiceImpl ...
type MediaServiceImpl struct {
	markup.Component `id:"MediaService"`

	AC                 application.Context           `inject:"context"`
	MediaDAO           dao.MediaDAO                  `inject:"#MediaDAO"`
	SysMainRepoService service.MainRepositoryService `inject:"#MainRepositoryService"`
	FileSystemService  service.FileSystemService     `inject:"#FileSystemService"`
	ContentTypeService service.ContentTypeService    `inject:"#ContentTypeService"`

	ResPathPrefix string `inject:"${wpm.presets.res-path-prefix}"`
	WebPathPrefix string `inject:"${wpm.presets.web-path-prefix}"`
}

func (inst *MediaServiceImpl) _Impl() service.MediaService {
	return inst
}

func (inst *MediaServiceImpl) dto2entity(o1 *dto.Media) (*entity.Media, error) {
	o2 := &entity.Media{}
	o2.ID = o1.ID

	o2.Bucket = o1.Bucket
	o2.ContentType = o1.ContentType
	o2.FileSize = o1.FileSize
	o2.Label = o1.Label
	o2.Name = o1.Name
	o2.SHA256SUM = o1.SHA256SUM
	o2.URL = o1.URL

	return o2, nil
}

func (inst *MediaServiceImpl) entity2dto(o1 *entity.Media, opt *service.MediaOptions) (*dto.Media, error) {

	if opt == nil {
		opt = &service.MediaOptions{}
	}

	o2 := &dto.Media{}
	o2.ID = o1.ID

	o2.Bucket = o1.Bucket
	o2.ContentType = o1.ContentType
	o2.FileSize = o1.FileSize
	o2.Label = o1.Label
	o2.Name = o1.Name
	o2.SHA256SUM = o1.SHA256SUM
	o2.URL = o1.URL

	return o2, nil
}

func (inst *MediaServiceImpl) entity2dtoList(list1 []*entity.Media, opt *service.MediaOptions) ([]*dto.Media, error) {

	if opt == nil {
		opt = &service.MediaOptions{}
	}

	list2 := make([]*dto.Media, 0)
	for _, it1 := range list1 {
		it2, err := inst.entity2dto(it1, opt)
		if err == nil {
			list2 = append(list2, it2)
		} else {
			vlog.Warn(err)
		}
	}
	return list2, nil
}

// ListAll ...
func (inst *MediaServiceImpl) ListAll(ctx context.Context, opt *service.MediaOptions) ([]*dto.Media, error) {
	list1, err := inst.MediaDAO.ListAll()
	if err != nil {
		return nil, err
	}
	return inst.entity2dtoList(list1, opt)
}

// Find ...
func (inst *MediaServiceImpl) Find(ctx context.Context, id dxo.MediaID, opt *service.MediaOptions) (*dto.Media, error) {
	o1, err := inst.MediaDAO.Find(id)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o1, opt)
}

// FindByIDs ...
func (inst *MediaServiceImpl) FindByIDs(ctx context.Context, ids []dxo.MediaID, opt *service.MediaOptions) ([]*dto.Media, error) {
	list1, err := inst.MediaDAO.ListByIDs(ids)
	if err != nil {
		return nil, err
	}
	return inst.entity2dtoList(list1, opt)
}

// FindByPath ...
func (inst *MediaServiceImpl) FindByPath(ctx context.Context, path string, opt *service.MediaOptions) (*dto.Media, error) {
	o1, err := inst.MediaDAO.FindByPath(path)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o1, opt)
}

func (inst *MediaServiceImpl) prepareInsert(ctx context.Context, o1 *dto.Media) (*entity.Media, error) {

	mdir, err := inst.getSystemMediaDir(ctx)
	if err != nil {
		return nil, err
	}

	src := inst.FileSystemService.Path(o1.LocalFilePath)
	dst, err := makeMediaLocalPath(o1, mdir)
	if err != nil {
		return nil, err
	}

	url, err := makeMediaWebPath(o1)
	if err != nil {
		return nil, err
	}

	sum, err := utils.ComputeFileSHA256sumAFS(src)
	if err != nil {
		return nil, err
	}

	o2 := &dto.Media{}
	o2.LocalFilePath = dst.GetPath()
	o2.URL = url
	o2.FileSize = src.GetInfo().Length()
	o2.ContentType = o1.ContentType
	o2.Name = o1.Name
	o2.Label = o1.Label
	o2.Bucket = "default"
	o2.SHA256SUM = sum

	err = src.CopyTo(dst, &afs.Options{
		Mkdirs: true, Create: true,
	})
	if err != nil {
		return nil, err
	}

	return inst.dto2entity(o2)
}

// Insert ...
func (inst *MediaServiceImpl) Insert(ctx context.Context, o *dto.Media) (*dto.Media, error) {
	o2, err := inst.prepareInsert(ctx, o)
	if err != nil {
		return nil, err
	}
	o3, err := inst.MediaDAO.Insert(o2)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o3, nil)
}

// ImportPresets ...
func (inst *MediaServiceImpl) ImportPresets(ctx context.Context) error {
	loader := myImportPresetsLoader{
		context: ctx, parent: inst,
	}
	return loader.Load()
}

// Update ...
func (inst *MediaServiceImpl) Update(ctx context.Context, id dxo.MediaID, o *dto.Media) (*dto.Media, error) {
	o2, err := inst.dto2entity(o)
	if err != nil {
		return nil, err
	}
	o3, err := inst.MediaDAO.Update(id, o2)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o3, nil)
}

// Remove ...
func (inst *MediaServiceImpl) Remove(ctx context.Context, id dxo.MediaID) error {
	return inst.MediaDAO.Remove(id)
}

// PrepareForDownload ...
func (inst *MediaServiceImpl) PrepareForDownload(ctx context.Context, me *dto.Media) (*dto.Media, error) {

	mediaDir, err := inst.getSystemMediaDir(ctx)
	if err != nil {
		return nil, err
	}

	pathDst, err := makeMediaLocalPath(me, mediaDir)
	if err != nil {
		return nil, err
	}

	me.LocalFilePath = pathDst.GetPath()
	return me, nil
}

func (inst *MediaServiceImpl) getSystemMediaDir(ctx context.Context) (afs.Path, error) {
	repo, err := inst.SysMainRepoService.GetRepository(ctx)
	if err != nil {
		return nil, err
	}
	objs := repo.Layout().Objects()
	p := objs.GetParent().GetChild("files/media")
	return p, nil
}

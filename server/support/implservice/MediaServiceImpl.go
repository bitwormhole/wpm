package implservice

import (
	"context"
	"errors"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// MediaServiceImpl ...
type MediaServiceImpl struct {
	markup.Component `id:"MediaService"`

	MediaDAO           dao.MediaDAO                  `inject:"#MediaDAO"`
	SysMainRepoService service.MainRepositoryService `inject:"#MainRepositoryService"`
}

func (inst *MediaServiceImpl) _Impl() service.MediaService {
	return inst
}

func (inst *MediaServiceImpl) dto2entity(o1 *dto.Media) (*entity.Media, error) {
	o2 := &entity.Media{}
	o2.ID = o1.ID

	o2.ContentType = o1.ContentType
	o2.FileSize = o1.FileSize
	o2.Path = o1.Path
	o2.SHA256SUM = o1.SHA256SUM

	return o2, nil
}

func (inst *MediaServiceImpl) entity2dto(o1 *entity.Media) (*dto.Media, error) {
	o2 := &dto.Media{}
	o2.ID = o1.ID

	o2.ContentType = o1.ContentType
	o2.FileSize = o1.FileSize
	o2.Path = o1.Path
	o2.SHA256SUM = o1.SHA256SUM

	return o2, nil
}

// ListAll ...
func (inst *MediaServiceImpl) ListAll(ctx context.Context) ([]*dto.Media, error) {

	// src, err := inst.MediaDAO.ListAll()
	// if err != nil {
	// 	return nil, err
	// }
	// dst := make([]*dto.Media, 0)
	// for _, item1 := range src {
	// 	item2, err := inst.entity2dto(item1)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	dst = append(dst, item2)
	// }
	// return dst, nil

	return nil, errors.New("no impl")
}

// Find ...
func (inst *MediaServiceImpl) Find(ctx context.Context, id dxo.MediaID) (*dto.Media, error) {
	return nil, errors.New("no impl")
}

// Insert ...
func (inst *MediaServiceImpl) Insert(ctx context.Context, o *dto.Media) (*dto.Media, error) {
	return nil, errors.New("no impl")
}

// Update ...
func (inst *MediaServiceImpl) Update(ctx context.Context, id dxo.MediaID, o *dto.Media) (*dto.Media, error) {
	return nil, errors.New("no impl")
}

// Remove ...
func (inst *MediaServiceImpl) Remove(ctx context.Context, id dxo.MediaID) error {
	return errors.New("no impl")
}

// FindByPath ...
func (inst *MediaServiceImpl) FindByPath(ctx context.Context, path string) (*dto.Media, error) {
	o1, err := inst.MediaDAO.FindByPath(path)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o1)
}

// PrepareForDownload ...
func (inst *MediaServiceImpl) PrepareForDownload(ctx context.Context, me *dto.Media) (*dto.Media, error) {

	// pathLF := me.LocalFilePath

	mediaDir, err := inst.getSystemMediaDir(ctx)
	if err != nil {
		return nil, err
	}

	pathDst := mediaDir.GetChild("/" + me.Path)

	// if !pathDst.Exists() {
	// }

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

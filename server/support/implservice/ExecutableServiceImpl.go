package implservice

import (
	"context"
	"errors"

	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ExecutableServiceImpl ...
type ExecutableServiceImpl struct {
	markup.Component `id:"ExecutableService"`

	ExecutableDAO     dao.ExecutableDAO         `inject:"#ExecutableDAO"`
	IconService       service.AppIconService    `inject:"#AppIconService"`
	FileSystemService service.FileSystemService `inject:"#FileSystemService"`
}

func (inst *ExecutableServiceImpl) _Impl() service.ExecutableService {
	return inst
}

func (inst *ExecutableServiceImpl) dto2entity(o1 *dto.Executable) (*entity.Executable, error) {
	o2 := &entity.Executable{}
	o2.ID = o1.ID
	o2.Name = o1.Name
	o2.Title = o1.Title
	o2.Path = o1.Path
	o2.Size = o1.Size
	o2.SHA256SUM = o1.SHA256SUM
	o2.Description = o1.Description
	o2.IconURL = o1.IconURL
	o2.OpenWithPriority = o1.OpenWithPriority
	// todo ...
	return o2, nil
}

func (inst *ExecutableServiceImpl) entity2dto(o1 *entity.Executable) (*dto.Executable, error) {
	o2 := &dto.Executable{}
	o2.ID = o1.ID
	o2.Name = o1.Name
	o2.Title = o1.Title
	o2.Path = o1.Path
	o2.Size = o1.Size
	o2.SHA256SUM = o1.SHA256SUM
	o2.Description = o1.Description
	o2.IconURL = o1.IconURL
	o2.OpenWithPriority = o1.OpenWithPriority

	// todo ...
	o2.State = inst.checkExeFileState(o1)

	// inst.IconService.FillWithIconURL(o2)

	return o2, nil
}

func (inst *ExecutableServiceImpl) checkExeFileState(o1 *entity.Executable) dxo.FileState {
	path := inst.FileSystemService.Path(o1.Path)
	if path.IsFile() {
		return dxo.FileStateReady
	} else if path.IsDirectory() {
		return dxo.FileStateUnknowError
	}
	return dxo.FileStateOffline
}

func (inst *ExecutableServiceImpl) checkBeforeInsert(ctx context.Context, o *dto.Executable) error {
	path := fs.Default().GetPath(o.Path)
	if !path.IsFile() {
		return errors.New("the file is not a executable, path=" + o.Path)
	}
	old, _ := inst.FindByPath(ctx, o.Path)
	if old != nil {
		return errors.New("the executable is exists, path=" + o.Path)
	}
	return nil
}

// ListAll ...
func (inst *ExecutableServiceImpl) ListAll(ctx context.Context) ([]*dto.Executable, error) {
	src, err := inst.ExecutableDAO.ListAll()
	if err != nil {
		return nil, err
	}
	dst := make([]*dto.Executable, 0)
	for _, item1 := range src {
		item2, err := inst.entity2dto(item1)
		if err != nil {
			return nil, err
		}
		dst = append(dst, item2)
	}
	return dst, nil
}

// Find ...
func (inst *ExecutableServiceImpl) Find(ctx context.Context, id dxo.ExecutableID) (*dto.Executable, error) {
	o1, err := inst.ExecutableDAO.Find(id)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o1)
}

// Insert ...
func (inst *ExecutableServiceImpl) Insert(ctx context.Context, o1 *dto.Executable) (*dto.Executable, error) {
	err := inst.checkBeforeInsert(ctx, o1)
	if err != nil {
		return nil, err
	}
	e1, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}
	e2, err := inst.ExecutableDAO.Insert(e1)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(e2)
}

// Update ...
func (inst *ExecutableServiceImpl) Update(ctx context.Context, id dxo.ExecutableID, o1 *dto.Executable) (*dto.Executable, error) {
	o2, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}
	o3, err := inst.ExecutableDAO.Update(id, o2)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o3)
}

// Remove ...
func (inst *ExecutableServiceImpl) Remove(ctx context.Context, id dxo.ExecutableID) error {
	return inst.ExecutableDAO.Remove(id)
}

// FindByPath ...
func (inst *ExecutableServiceImpl) FindByPath(ctx context.Context, path string) (*dto.Executable, error) {
	e1, err := inst.ExecutableDAO.FindByPath(path)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(e1)
}

package locations

import (
	"context"
	"fmt"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ImpLocationService ...
type ImpLocationService struct {
	markup.Component `id:"LocationService"`

	DAO               dao.LocationDAO           `inject:"#LocationDAO"`
	FileSystemService service.FileSystemService `inject:"#FileSystemService"`
}

func (inst *ImpLocationService) _Impl() service.LocationService {
	return inst
}

func (inst *ImpLocationService) dto2entity(o1 *dto.Location) (*entity.Location, error) {

	o2 := &entity.Location{}

	o2.ID = o1.ID
	o2.UUID = o1.UUID
	o2.AsDir = o1.AsDir
	o2.AsFile = o1.AsFile
	o2.Path = o1.Path
	o2.Class = o1.Class
	o2.Target = o1.Target

	// o2.TypeID = o1.TypeID
	// o2.TypeKey = o1.TypeKey
	// o2.TypeName = o1.TypeName

	return o2, nil
}

func (inst *ImpLocationService) entity2dto(o1 *entity.Location, opt *service.LocationOptions) (*dto.Location, error) {

	if opt == nil {
		opt = &service.LocationOptions{}
	}

	o2 := &dto.Location{}
	o2.ID = o1.ID
	o2.UUID = o1.UUID
	o2.CreatedAt = util.NewTime(o1.CreatedAt)
	o2.UpdatedAt = util.NewTime(o1.UpdatedAt)

	o2.AsDir = o1.AsDir
	o2.AsFile = o1.AsFile
	o2.CreatedAt = util.NewTime(o1.CreatedAt)
	o2.UpdatedAt = util.NewTime(o1.UpdatedAt)
	o2.Path = o1.Path
	o2.Class = o1.Class
	o2.Target = o1.Target

	// o2.TypeID = o1.TypeID
	// o2.TypeKey = o1.TypeKey
	// o2.TypeName = o1.TypeName

	return o2, nil
}

func (inst *ImpLocationService) entity2dtoList(src []*entity.Location, opt *service.LocationOptions) ([]*dto.Location, error) {

	if opt == nil {
		opt = &service.LocationOptions{}
	}

	dst := make([]*dto.Location, 0)
	for _, item1 := range src {
		item2, err := inst.entity2dto(item1, opt)
		if err != nil {
			return nil, err
		}
		dst = append(dst, item2)
	}
	return dst, nil
}

// Find ...
func (inst *ImpLocationService) Find(ctx context.Context, id dxo.LocationID, opt *service.LocationOptions) (*dto.Location, error) {
	o1, err := inst.DAO.Find(id)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o1, opt)
}

// FindByPath ...
func (inst *ImpLocationService) FindByPath(ctx context.Context, path string, opt *service.LocationOptions) (*dto.Location, error) {
	o1, err := inst.DAO.FindByPath(path)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o1, opt)
}

// ListAll ...
func (inst *ImpLocationService) ListAll(ctx context.Context, opt *service.LocationOptions) ([]*dto.Location, error) {
	list, err := inst.DAO.ListAll()
	if err != nil {
		return nil, err
	}
	return inst.entity2dtoList(list, opt)
}

func (inst *ImpLocationService) checkBeforeInsert(ctx context.Context, l *entity.Location) error {

	path := l.Path
	node := inst.FileSystemService.Path(path)

	if l.AsDir {
		if node.IsDirectory() {
			return nil
		}
	}

	if l.AsFile {
		if node.IsFile() {
			return nil
		}
	}

	return fmt.Errorf("no file (or dir): %v", path)
}

// Insert ...
func (inst *ImpLocationService) Insert(ctx context.Context, o1 *dto.Location) (*dto.Location, error) {

	o2, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}

	err = inst.checkBeforeInsert(ctx, o2)
	if err != nil {
		return nil, err
	}

	o3, err := inst.DAO.Insert(o2)
	if err != nil {
		return nil, err
	}

	return inst.entity2dto(o3, nil)
}

// Update ...
func (inst *ImpLocationService) Update(ctx context.Context, id dxo.LocationID, o1 *dto.Location) (*dto.Location, error) {
	o2, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}
	o3, err := inst.DAO.Update(id, o2)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o3, nil)
}

// Remove ...
func (inst *ImpLocationService) Remove(ctx context.Context, id dxo.LocationID) error {
	return inst.DAO.Remove(id)
}

// InsertOrFetch ...
func (inst *ImpLocationService) InsertOrFetch(ctx context.Context, o1 *dto.Location, opt *service.LocationOptions) (*dto.Location, error) {
	// try find
	path := o1.Path
	o2, err := inst.FindByPath(ctx, path, opt)
	if err == nil && o2 != nil {
		return o2, nil
	}
	// do insert
	return inst.Insert(ctx, o1)
}

// PutExecutable ...
func (inst *ImpLocationService) PutExecutable(ctx context.Context, o *entity.Executable) error {
	o1 := &dto.Location{
		Path:   o.Path,
		Target: dxo.URN(o.URN),
		Class:  dxo.LocationExecutable,
		AsFile: true,
		AsDir:  false,
	}
	_, err := inst.Insert(ctx, o1)
	return err
}

// PutProject ...
func (inst *ImpLocationService) PutProject(ctx context.Context, o *entity.Project) error {
	o1 := &dto.Location{
		Path:   o.RegularPath,
		Target: o.URN,
		Class:  dxo.LocationProject,
		AsFile: true,
		AsDir:  true,
	}
	_, err := inst.Insert(ctx, o1)
	return err
}

// PutRepository ...
func (inst *ImpLocationService) PutRepository(ctx context.Context, o *entity.LocalRepository) error {
	o1 := &dto.Location{
		Path:   o.RegularPath,
		Target: o.URN,
		Class:  dxo.LocationGitRepository,
		AsFile: true,
		AsDir:  false,
	}
	_, err := inst.Insert(ctx, o1)
	return err
}

// PutWorktree ...
func (inst *ImpLocationService) PutWorktree(ctx context.Context, o *entity.Worktree) error {
	o1 := &dto.Location{
		Path:   o.RegularPath,
		Target: o.URN,
		Class:  dxo.LocationGitWorktree,
		AsFile: true,
		AsDir:  true,
	}
	_, err := inst.Insert(ctx, o1)
	return err
}

package worktrees

import (
	"context"
	"strings"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ImpWorktreeService ...
type ImpWorktreeService struct {
	markup.Component `id:"WorktreeService"`

	FileSystemService service.FileSystemService `inject:"#FileSystemService"`
	DAO               dao.WorktreeDAO           `inject:"#WorktreeDAO"`
}

func (inst *ImpWorktreeService) _Impl() service.WorktreeService {
	return inst
}

func (inst *ImpWorktreeService) entity2dto(c context.Context, o1 *entity.Worktree, opt *service.WorktreeOptions) (*dto.Worktree, error) {

	if opt == nil {
		opt = &service.WorktreeOptions{}
	}

	o2 := &dto.Worktree{}

	o2.ID = o1.ID
	o2.Name = o1.Name
	o2.UUID = o1.UUID
	o2.OwnerRepository = o1.OwnerRepository

	o2.WorkingDir = o1.WorkingDirectory
	o2.DotGitPath = o1.DotGitPath

	if opt.WithFileState {
		o2.State = inst.checkFileState(o1)
	}

	return o2, nil
}

func (inst *ImpWorktreeService) dto2entity(o1 *dto.Worktree) (*entity.Worktree, error) {

	o2 := &entity.Worktree{}

	o2.ID = o1.ID
	o2.Name = o1.Name
	o2.UUID = o1.UUID
	o2.OwnerRepository = o1.OwnerRepository

	o2.WorkingDirectory = o1.WorkingDir
	o2.DotGitPath = o1.DotGitPath

	return o2, nil
}

func (inst *ImpWorktreeService) entity2dtoList(c context.Context, src []*entity.Worktree, opt *service.WorktreeOptions) ([]*dto.Worktree, error) {

	if opt == nil {
		opt = &service.WorktreeOptions{}
	}

	dst := make([]*dto.Worktree, 0)
	for _, item1 := range src {
		item2, err := inst.entity2dto(c, item1, opt)
		if err != nil {
			item2 = nil
		}
		if item2 != nil {
			dst = append(dst, item2)
		}
	}
	return dst, nil
}

func (inst *ImpWorktreeService) checkFileState(o1 *entity.Worktree) dxo.FileState {
	path := inst.FileSystemService.Path(o1.DotGitPath)
	if path.Exists() {
		return dxo.FileStateReady
	}
	return dxo.FileStateOffline
}

// Find ...
func (inst *ImpWorktreeService) Find(ctx context.Context, id dxo.WorktreeID, opt *service.WorktreeOptions) (*dto.Worktree, error) {
	o1, err := inst.DAO.Find(id)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(ctx, o1, opt)
}

// FindByPath ...
func (inst *ImpWorktreeService) FindByPath(ctx context.Context, path string, opt *service.WorktreeOptions) (*dto.Worktree, error) {
	o1, err := inst.DAO.FindByPath(path)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(ctx, o1, opt)
}

// ListAll ...
func (inst *ImpWorktreeService) ListAll(ctx context.Context, opt *service.WorktreeOptions) ([]*dto.Worktree, error) {
	list, err := inst.DAO.ListAll()
	if err != nil {
		return nil, err
	}
	return inst.entity2dtoList(ctx, list, opt)
}

// Insert ...
func (inst *ImpWorktreeService) Insert(ctx context.Context, o1 *dto.Worktree) (*dto.Worktree, error) {
	o2, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}
	o3, err := inst.DAO.Insert(o2)
	if err != nil {
		return nil, err
	}
	opt := &service.WorktreeOptions{WithFileState: true}
	return inst.entity2dto(ctx, o3, opt)
}

// InsertOrFetch ...
func (inst *ImpWorktreeService) InsertOrFetch(ctx context.Context, o1 *dto.Worktree, opt *service.WorktreeOptions) (*dto.Worktree, error) {

	o2, err := inst.Insert(ctx, o1)
	if o2 != nil && err == nil {
		return o2, nil
	}

	errlist := &utils.ErrorList{}
	errlist.Append(err)

	paths := make([]string, 0)
	paths = append(paths, o1.DotGitPath)
	paths = append(paths, o1.WorkingDir)

	for _, wantPath := range paths {
		wantPath = strings.TrimSpace(wantPath)
		if wantPath == "" {
			continue
		}
		o2, err := inst.FindByPath(ctx, wantPath, opt)
		if o2 != nil && err == nil {
			return o2, nil
		}
		errlist.Append(err)
	}

	return nil, errlist.First()
}

// Update ...
func (inst *ImpWorktreeService) Update(ctx context.Context, id dxo.WorktreeID, o1 *dto.Worktree) (*dto.Worktree, error) {
	o2, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}
	o3, err := inst.DAO.Update(id, o2)
	if err != nil {
		return nil, err
	}
	opt := &service.WorktreeOptions{WithFileState: true}
	return inst.entity2dto(ctx, o3, opt)
}

// Remove ...
func (inst *ImpWorktreeService) Remove(ctx context.Context, id dxo.WorktreeID) error {
	return inst.DAO.Remove(id)
}

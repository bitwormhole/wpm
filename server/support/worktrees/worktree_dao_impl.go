package worktrees

import (
	"fmt"
	"strings"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dbagent"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
)

// ImpWorktreeDao ...
type ImpWorktreeDao struct {
	markup.Component `id:"WorktreeDAO"`

	Agent          dbagent.GormDBAgent    `inject:"#GormDBAgent"`
	TrashService   service.TrashService   `inject:"#TrashService"`
	UUIDGenService service.UUIDGenService `inject:"#UUIDGenService"`
}

func (inst *ImpWorktreeDao) _Impl() dao.WorktreeDAO {
	return inst
}

func (inst *ImpWorktreeDao) model() *entity.Worktree {
	return &entity.Worktree{}
}

func (inst *ImpWorktreeDao) modelList() []*entity.Worktree {
	return make([]*entity.Worktree, 0)
}

// Find ...
func (inst *ImpWorktreeDao) Find(id dxo.WorktreeID) (*entity.Worktree, error) {
	db := inst.Agent.DB()
	o2 := inst.model()
	res := db.First(o2, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return o2, nil
}

// FindByPath ...
func (inst *ImpWorktreeDao) FindByPath(path string) (*entity.Worktree, error) {

	path = strings.TrimSpace(path)
	if path == "" {
		return nil, fmt.Errorf("param:`path` is empty")
	}

	erlist := &utils.ErrorList{}
	db := inst.Agent.DB()
	o := inst.model()
	fields := o.ListPathFields()

	for _, field := range fields {
		res := db.Where(field+"=?", path).First(o)
		if res.Error == nil {
			return o, nil
		}
		erlist.Append(res.Error)
	}

	return nil, erlist.First()
}

// ListAll ...
func (inst *ImpWorktreeDao) ListAll() ([]*entity.Worktree, error) {
	list := inst.modelList()
	db := inst.Agent.DB()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

// Insert ...
func (inst *ImpWorktreeDao) Insert(o *entity.Worktree) (*entity.Worktree, error) {

	inst.TrashService.OnInsert()

	// compute fields
	o.ID = 0
	o.UUID = inst.UUIDGenService.GenerateUUID(o.DotGitPath + "|entity.Project|" + o.WorkingDirectory)

	// save
	db := inst.Agent.DB()
	res := db.Create(o)
	if res.Error != nil {
		return nil, res.Error
	}

	return o, nil
}

// Update ...
func (inst *ImpWorktreeDao) Update(id dxo.WorktreeID, o1 *entity.Worktree) (*entity.Worktree, error) {

	db := inst.Agent.DB()
	o2 := inst.model()
	res := db.First(o2, id)
	if res.Error != nil {
		return nil, res.Error
	}

	// todo ...
	o2.Name = o1.Name
	o2.WorkingDirectory = o1.WorkingDirectory
	o2.DotGitPath = o1.DotGitPath
	o2.OwnerRepository = o1.OwnerRepository

	res = db.Save(o2)
	if res.Error != nil {
		return nil, res.Error
	}
	return o2, nil
}

// Remove ...
func (inst *ImpWorktreeDao) Remove(id dxo.WorktreeID) error {

	inst.TrashService.OnDelete()

	db := inst.Agent.DB()
	m := inst.model()
	res := db.Delete(m, id)
	return res.Error
}

package repositories

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dbagent"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
)

// RepositoryDaoImpl ...
type RepositoryDaoImpl struct {
	markup.Component `id:"LocalRepositoryDAO"`

	Agent          dbagent.GormDBAgent    `inject:"#GormDBAgent"`
	TrashService   service.TrashService   `inject:"#TrashService"`
	UUIDGenService service.UUIDGenService `inject:"#UUIDGenService"`
}

func (inst *RepositoryDaoImpl) _Impl() dao.LocalRepositoryDAO {
	return inst
}

func (inst *RepositoryDaoImpl) model() *entity.LocalRepository {
	return &entity.LocalRepository{}
}

func (inst *RepositoryDaoImpl) modelList() []*entity.LocalRepository {
	return make([]*entity.LocalRepository, 0)
}

func (inst *RepositoryDaoImpl) normalizePath(path string) string {
	deffs := fs.Default()
	p2 := deffs.GetPath(path)
	if p2 == nil {
		return ""
	}
	return p2.String()
}

func (inst *RepositoryDaoImpl) checkEntity(o *entity.LocalRepository) error {
	path := o.Path
	o2, err := inst.FindByPath(path)
	if err == nil && o2 != nil {
		if o2.ID != o.ID {
			return errors.New("the path is duplicate: " + path)
		}
	}
	return nil
}

// Find ...
func (inst *RepositoryDaoImpl) Find(id dxo.LocalRepositoryID) (*entity.LocalRepository, error) {
	db := inst.Agent.DB()
	o := inst.model()
	res := db.First(o, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// FindByUUID ...
func (inst *RepositoryDaoImpl) FindByUUID(uuid dxo.UUID) (*entity.LocalRepository, error) {
	db := inst.Agent.DB()
	o := inst.model()
	res := db.Where("uuid=?", uuid).First(o)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// FindByName ...
func (inst *RepositoryDaoImpl) FindByName(name string) (*entity.LocalRepository, error) {
	db := inst.Agent.DB()
	o := inst.model()
	res := db.Where("name=?", name).First(o)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// FindByPath ...
func (inst *RepositoryDaoImpl) FindByPath(path string) (*entity.LocalRepository, error) {

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

// FindByDotGit ...
func (inst *RepositoryDaoImpl) FindByDotGit(path string) (*entity.LocalRepository, error) {
	db := inst.Agent.DB()
	o := inst.model()
	res := db.Where("dot_git_path=?", path).First(o)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// FindByWorkingDir ...
func (inst *RepositoryDaoImpl) FindByWorkingDir(path string) (*entity.LocalRepository, error) {
	db := inst.Agent.DB()
	o := inst.model()
	res := db.Where("working_path=?", path).First(o)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// ListAll ...
func (inst *RepositoryDaoImpl) ListAll() ([]*entity.LocalRepository, error) {
	db := inst.Agent.DB()
	list := inst.modelList()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

// ListByIds ...
func (inst *RepositoryDaoImpl) ListByIds(ids []dxo.LocalRepositoryID) ([]*entity.LocalRepository, error) {

	if ids == nil {
		return []*entity.LocalRepository{}, nil
	}

	if len(ids) == 0 {
		return []*entity.LocalRepository{}, nil
	}

	db := inst.Agent.DB()
	list := inst.modelList()
	res := db.Find(&list, ids)
	if res.Error != nil {
		return nil, res.Error
	}

	return list, nil
}

// Insert ...
func (inst *RepositoryDaoImpl) Insert(o *entity.LocalRepository) (*entity.LocalRepository, error) {

	inst.TrashService.OnInsert()

	db := inst.Agent.DB()

	// compute fields
	o.ID = 0
	o.UUID = inst.UUIDGenService.GenerateUUID(o.Path + "|entity.LocalRepository|" + o.DotGitPath)
	o.Base.PrepareInsert()

	o.Path = inst.normalizePath(o.Path)
	o.DotGitPath = inst.normalizePath(o.DotGitPath)

	// check
	err := inst.checkEntity(o)
	if err != nil {
		return nil, err
	}

	// save
	res := db.Create(o)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// Update ...
func (inst *RepositoryDaoImpl) Update(id dxo.LocalRepositoryID, o1 *entity.LocalRepository) (*entity.LocalRepository, error) {
	db := inst.Agent.DB()
	o2 := inst.model()
	res := db.First(o2, id)
	if res.Error != nil {
		return nil, res.Error
	}

	// compute fields
	o2.Path = inst.normalizePath(o1.Path)
	o2.ConfigFile = inst.normalizePath(o1.ConfigFile)
	o2.DotGitPath = inst.normalizePath(o1.DotGitPath)
	o2.WorkingPath = inst.normalizePath(o1.WorkingPath)
	o2.RepositoryPath = inst.normalizePath(o1.RepositoryPath)

	o2.Name = o1.Name
	o2.DisplayName = o1.DisplayName
	o2.Description = o1.Description

	// check
	err := inst.checkEntity(o2)
	if err != nil {
		return nil, err
	}

	// save
	res = db.Save(o2)
	if res.Error != nil {
		return nil, res.Error
	}
	return o2, nil
}

// Remove ...
func (inst *RepositoryDaoImpl) Remove(id dxo.LocalRepositoryID) error {

	inst.TrashService.OnDelete()

	db := inst.Agent.DB()
	o := inst.model()
	o.ID = id
	res := db.Delete(o, id)
	return res.Error
}

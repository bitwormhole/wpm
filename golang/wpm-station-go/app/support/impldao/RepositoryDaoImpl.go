package impldao

import (
	"errors"

	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/app/data/dao"
	"github.com/bitwormhole/wpm/app/data/dxo"
	"github.com/bitwormhole/wpm/app/data/entity"
	"github.com/bitwormhole/wpm/app/service"
)

// RepositoryDaoImpl ...
type RepositoryDaoImpl struct {
	markup.Component `id:"RepositoryDAO"`

	Agent          GormDBAgent            `inject:"#GormDBAgent"`
	UUIDGenService service.UUIDGenService `inject:"#UUIDGenService"`
}

func (inst *RepositoryDaoImpl) _Impl() dao.RepositoryDAO {
	return inst
}

func (inst *RepositoryDaoImpl) model() *entity.Repository {
	return &entity.Repository{}
}

func (inst *RepositoryDaoImpl) modelList() []*entity.Repository {
	return make([]*entity.Repository, 0)
}

func (inst *RepositoryDaoImpl) normalizePath(path string) string {
	deffs := fs.Default()
	p2 := deffs.GetPath(path)
	if p2 == nil {
		return ""
	}
	return p2.String()
}

func (inst *RepositoryDaoImpl) checkEntity(o *entity.Repository) error {
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
func (inst *RepositoryDaoImpl) Find(id dxo.RepositoryID) (*entity.Repository, error) {
	db := inst.Agent.DB()
	o := inst.model()
	res := db.First(o, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// FindByName ...
func (inst *RepositoryDaoImpl) FindByName(name string) (*entity.Repository, error) {
	db := inst.Agent.DB()
	o := inst.model()
	res := db.Where("name=?", name).First(o)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// FindByPath ...
func (inst *RepositoryDaoImpl) FindByPath(path string) (*entity.Repository, error) {
	db := inst.Agent.DB()
	o := inst.model()
	res := db.Where("path=?", path).First(o)
	if res.Error != nil {
		return nil, res.Error
	}
	return o, nil
}

// ListAll ...
func (inst *RepositoryDaoImpl) ListAll() ([]*entity.Repository, error) {
	db := inst.Agent.DB()
	list := inst.modelList()
	res := db.Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}

// Insert ...
func (inst *RepositoryDaoImpl) Insert(o *entity.Repository) (*entity.Repository, error) {
	db := inst.Agent.DB()

	// compute fields
	o.ID = 0
	o.UUID = inst.UUIDGenService.GenerateUUID(o.Path + "|entity.Repository|" + o.DotGitPath)
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
func (inst *RepositoryDaoImpl) Update(id dxo.RepositoryID, o1 *entity.Repository) (*entity.Repository, error) {
	db := inst.Agent.DB()
	o2 := inst.model()
	res := db.First(o2, id)
	if res.Error != nil {
		return nil, res.Error
	}

	// compute fields
	o2.Path = inst.normalizePath(o1.Path)
	o2.DotGitPath = inst.normalizePath(o1.DotGitPath)

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
func (inst *RepositoryDaoImpl) Remove(id dxo.RepositoryID) error {
	db := inst.Agent.DB()
	o := inst.model()
	res := db.Delete(o, id)
	return res.Error
}

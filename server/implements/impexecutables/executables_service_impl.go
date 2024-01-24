package impexecutables

import (
	"context"
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/classes/executables"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/starter-go/afs"
	"github.com/starter-go/base/lang"
)

// ServiceImpl ...
type ServiceImpl struct {

	//starter:component
	_as func(executables.Service) //starter:as("#")

	Dao         executables.DAO         //starter:inject("#")
	StateLoader executables.StateLoader //starter:inject("#")
	FS          afs.FS                  //starter:inject("#")

	convertor executables.Convertor
}

func (inst *ServiceImpl) _impl() executables.Service {
	return inst
}

func (inst *ServiceImpl) computeURN(obj *dto.Executable) dxo.ExecutableURN {
	// example: [urn:executable:default#explorer]
	ns := strings.TrimSpace(obj.Namespace)
	name := strings.TrimSpace(obj.Name)
	str := "urn:executable:" + ns + "#" + name
	str = strings.ToLower(str)
	return dxo.ExecutableURN(str)
}

func (inst *ServiceImpl) loadExeMeta(obj *dto.Executable) error {

	file := inst.FS.NewPath(obj.Path)
	if !file.IsFile() {
		path := file.GetPath()
		return fmt.Errorf("no file at path [%s]", path)
	}

	info := file.GetInfo()
	size := info.Length()
	bin, err := file.GetIO().ReadBinary(nil)
	if err != nil {
		return err
	}

	sum := sha256.Sum256(bin)
	hex := lang.HexFromBytes(sum[:])

	obj.SHA256SUM = hex
	obj.Size = size
	return nil
}

// Insert ...
func (inst *ServiceImpl) Insert(ctx context.Context, o1 *dto.Executable) (*dto.Executable, error) {

	o1.URN = inst.computeURN(o1)
	inst.loadExeMeta(o1)

	o2 := inst.convertor.ConvertD2E(o1)
	o3, err := inst.Dao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}
	o4 := inst.convertor.ConvertE2D(o3)
	return o4, nil
}

// Update ...
func (inst *ServiceImpl) Update(ctx context.Context, id dxo.ExecutableID, o1 *dto.Executable) (*dto.Executable, error) {

	inst.loadExeMeta(o1)

	o2, err := inst.Dao.Update(nil, id, func(ent *entity.Executable) {

		ent.Name = o1.Name
		ent.Aliases = o1.Aliases
		ent.Namespace = o1.Namespace
		ent.Title = o1.Title
		ent.Description = o1.Description
		ent.IconURL = o1.IconURL

		ent.Path = o1.Path
		ent.Size = o1.Size
		ent.SHA256SUM = o1.SHA256SUM
		ent.OpenWithPriority = o1.OpenWithPriority

		ent.OS = o1.OS
		ent.Arch = o1.Arch
		ent.Version = o1.Version

		if ent.URN == "" {
			ent.URN = inst.computeURN(o1)
		}

	})
	if err != nil {
		return nil, err
	}
	o3 := inst.convertor.ConvertE2D(o2)
	return o3, nil
}

// Remove ...
func (inst *ServiceImpl) Remove(ctx context.Context, id dxo.ExecutableID) error {
	return inst.Dao.Remove(nil, id)
}

// Find ...
func (inst *ServiceImpl) Find(ctx context.Context, id dxo.ExecutableID) (*dto.Executable, error) {

	o1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}

	o2 := inst.convertor.ConvertE2D(o1)
	err = inst.StateLoader.Load(ctx, id, o2)
	if err != nil {
		return nil, err
	}

	return o2, nil
}

// List ...
func (inst *ServiceImpl) List(ctx context.Context, q *executables.Query) ([]*dto.Executable, error) {
	list1, err := inst.Dao.List(nil, q)
	if err != nil {
		return nil, err
	}
	list2 := inst.convertor.ConvertListE2D(list1)

	for _, item := range list2 {
		err = inst.StateLoader.Load(ctx, item.ID, item)
		if err != nil {
			return nil, err
		}
	}

	return list2, nil
}

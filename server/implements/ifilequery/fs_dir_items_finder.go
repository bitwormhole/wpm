package ifilequery

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/vo"
	"github.com/starter-go/afs"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/vlog"
)

type fsDirItemsFinder struct {
	parent *ServiceImpl
}

func (inst *fsDirItemsFinder) find(want *vo.FileQuery) (*vo.FileQuery, error) {

	dst := new(vo.FileQuery)
	path, err := inst.pathOf(want)
	if err != nil {
		return nil, err
	}

	if path.IsFile() {
		err = inst.findAsFile(path, dst)
	} else if path.IsDirectory() {
		err = inst.findAsDir(path, dst)
	} else {
		p := path.GetPath()
		err = fmt.Errorf("bad path [%s]", p)
	}

	if err != nil {
		return nil, err
	}

	dst.BaseURL = "file:///"
	dst.Path = path.GetPath()
	dst.Self.Path = path.GetPath()
	return dst, nil
}

func (inst *fsDirItemsFinder) findAsDir(dir afs.Path, dst *vo.FileQuery) error {

	namelist := dir.ListNames()
	sort.Strings(namelist)

	for _, name := range namelist {
		child := dir.GetChild(name)
		item, err := inst.convertToDTO(child)
		if err == nil {
			dst.Items = append(dst.Items, item)
		} else {
			vlog.Warn(err.Error())
		}
	}

	self, err := inst.convertToDTO(dir)
	if err != nil {
		return err
	}
	dst.Self = *self
	return nil
}

func (inst *fsDirItemsFinder) findAsFile(file afs.Path, dst *vo.FileQuery) error {
	self, err := inst.convertToDTO(file)
	if err != nil {
		return err
	}
	dst.Self = *self
	return nil
}

func (inst *fsDirItemsFinder) convertToDTO(path afs.Path) (*dto.File, error) {

	info := path.GetInfo()
	createdAt := info.CreatedAt()
	updatedAt := info.UpdatedAt()

	item := new(dto.File)
	item.Name = path.GetName()
	item.Size = info.Length()
	item.Type = "todo..."
	item.Icon = "todo..."
	item.IsDir = info.IsDirectory()
	item.IsFile = info.IsFile()
	item.Exists = info.Exists()
	item.CreatedAt = lang.NewTime(createdAt)
	item.UpdatedAt = lang.NewTime(updatedAt)

	return item, nil
}

func (inst *fsDirItemsFinder) pathOf(want *vo.FileQuery) (afs.Path, error) {

	const prefix = "file:/"
	p1 := want.BaseURL
	p2 := want.Path

	if strings.HasPrefix(p1, prefix) {
		p1 = p1[len(prefix):]
	} else {
		return nil, fmt.Errorf("bad base URL [%s]", p1)
	}

	str := p1 + "/" + p2
	path := inst.parent.FS.NewPath(str)

	if !path.Exists() {
		str = path.GetPath()
		return nil, fmt.Errorf("file or folder is not exists, path = [%s]", str)
	}
	return path, nil
}

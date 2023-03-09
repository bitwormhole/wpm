package projecttypes

import (
	"context"
	"fmt"
	"strings"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/web/dto"
)

type myProjectLocatorWithTypes struct {

	// params
	context context.Context
	parent  *ProjectTypeServiceImpl
	project *dto.Project
	path    string

	// var
	types []*entity.ProjectType
}

func (inst *myProjectLocatorWithTypes) Locate() error {

	err := inst.checkParams()
	if err != nil {
		return err
	}

	err = inst.loadTypes()
	if err != nil {
		return err
	}

	err = inst.lookup()
	if err != nil {
		return err
	}

	return nil
}

func (inst *myProjectLocatorWithTypes) checkParams() error {
	if inst.parent == nil || inst.project == nil || inst.context == nil {
		return fmt.Errorf("param is nil")
	}
	if inst.path == "" {
		inst.path = inst.project.FullPath
	}
	if inst.path == "" {
		return fmt.Errorf("no param: 'path'")
	}
	return nil
}

func (inst *myProjectLocatorWithTypes) loadTypes() error {
	ptdao := inst.parent.ProjectTypeDAO
	tid := inst.project.ProjectType
	if tid == 0 {
		// load all
		all, err := ptdao.ListAll()
		if err != nil {
			return err
		}
		inst.types = all
	} else {
		// get one
		ent, err := ptdao.Find(tid)
		if err != nil {
			return err
		}
		inst.types = []*entity.ProjectType{ent}
	}
	return nil
}

func (inst *myProjectLocatorWithTypes) lookup() error {
	timeout := 99
	path := inst.parent.FileSystemService.Path(inst.path)
	for ; path != nil; path = path.GetParent() {
		if timeout > 0 {
			timeout--
		} else {
			return fmt.Errorf("the path is too deep, path=%v", inst.path)
		}
		for _, t := range inst.types {
			err := inst.tryLocateProject(path, t)
			if err == nil {
				return nil
			}
		}
	}
	return fmt.Errorf("no project in path [%v]", inst.path)
}

func (inst *myProjectLocatorWithTypes) tryLocateProject(dir afs.Path, pt *entity.ProjectType) error {

	if !dir.IsDirectory() {
		return fmt.Errorf("not a dir")
	}

	items := dir.ListChildren()
	for _, item := range items {
		if inst.isTypeMatch(item, pt) {
			return inst.apply(item, pt)
		}
	}

	return fmt.Errorf("no project")
}

func (inst *myProjectLocatorWithTypes) isTypeMatch(path afs.Path, pt *entity.ProjectType) bool {

	nameWant := pt.Name
	nameHave := path.GetName()
	if !inst.isFileNameMatch(nameWant, nameHave) {
		return false
	}

	if pt.AsDir && path.IsDirectory() {
		return true
	}

	if pt.AsFile && path.IsFile() {
		return true
	}

	return false
}

func (inst *myProjectLocatorWithTypes) isFileNameMatch(want, have string) bool {
	have = inst.normalizeFileName(have)
	want = inst.normalizeFileName(want) // maybe include wildcard like '*'

	// without wildcard
	if !strings.ContainsRune(want, '*') {
		return have == want
	}

	// with wildcard
	prefix := ""
	suffix := ""
	text := have
	parts := strings.Split(want, "*")
	for ipart, part := range parts {
		if ipart == 0 {
			prefix = part
		}
		suffix = part
		if part == "" {
			continue
		}
		i := strings.Index(text, part)
		if i < 0 {
			return false
		}
		text = text[i+len(part):]
	}
	if prefix != "" {
		if !strings.HasPrefix(have, prefix) {
			return false
		}
	}
	if suffix != "" {
		if !strings.HasSuffix(have, suffix) {
			return false
		}
	}

	return true
}

func (inst *myProjectLocatorWithTypes) normalizeFileName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return name
}

func (inst *myProjectLocatorWithTypes) apply(projectFile afs.Path, pt *entity.ProjectType) error {

	p := inst.project

	p.ProjectType = pt.ID
	p.ProjectTypeName = pt.Name
	p.ProjectDir = projectFile.GetParent().GetPath()
	p.FullPath = projectFile.GetPath()
	p.ConfigFileName = projectFile.GetName()

	if p.Name == "" {
		p.Name = projectFile.GetParent().GetName()
	}

	return nil
}

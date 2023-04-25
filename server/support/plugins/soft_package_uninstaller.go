package plugins

import (
	"context"
	"fmt"
	"strings"

	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/dbagent"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"gorm.io/gorm"
)

type myPakcageUninstaller struct {
	context      context.Context
	agent        dbagent.GormDBAgent
	installation dxo.InstallationID

	FileSystemSer service.FileSystemService

	db      *gorm.DB
	objects []any
	files   []*entity.InstalledFile
}

func (inst *myPakcageUninstaller) Uninstall() error {

	steps := make([]func() error, 0)
	steps = append(steps, inst.prepare)

	steps = append(steps, inst.findContentTypes)
	steps = append(steps, inst.findExecutables)
	steps = append(steps, inst.findIntentTemplates)
	steps = append(steps, inst.findMediae)
	steps = append(steps, inst.findInstalledFiles)

	steps = append(steps, inst.removeFiles)
	steps = append(steps, inst.removeObjects)

	for _, fn := range steps {
		err := fn()
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myPakcageUninstaller) prepare() error {

	iid := inst.installation
	if iid < 1000 {
		return fmt.Errorf("bad installation id")
	}

	inst.db = inst.agent.DB()
	return nil
}

func (inst *myPakcageUninstaller) findMediae() error {
	iid := inst.installation
	src := make([]*entity.Media, 0)
	res := inst.db.Where("installation = ?", iid).Find(&src)
	if res.Error != nil {
		return res.Error
	}
	for _, item := range src {
		inst.add(item)
	}
	return nil
}

func (inst *myPakcageUninstaller) findInstalledFiles() error {
	iid := inst.installation
	src := make([]*entity.InstalledFile, 0)
	res := inst.db.Where("installation = ?", iid).Find(&src)
	if res.Error != nil {
		return res.Error
	}
	for _, item := range src {
		inst.add(item)
		inst.addFile(item)
	}
	return nil
}

func (inst *myPakcageUninstaller) findContentTypes() error {
	iid := inst.installation
	src := make([]*entity.ContentType, 0)
	res := inst.db.Where("installation = ?", iid).Find(&src)
	if res.Error != nil {
		return res.Error
	}
	for _, item := range src {
		inst.add(item)
	}
	return nil
}

func (inst *myPakcageUninstaller) findIntentTemplates() error {
	iid := inst.installation
	src := make([]*entity.IntentTemplate, 0)
	res := inst.db.Where("installation = ?", iid).Find(&src)
	if res.Error != nil {
		return res.Error
	}
	for _, item := range src {
		inst.add(item)
	}
	return nil
}

func (inst *myPakcageUninstaller) findExecutables() error {
	iid := inst.installation
	src := make([]*entity.Executable, 0)
	res := inst.db.Where("installation = ?", iid).Find(&src)
	if res.Error != nil {
		return res.Error
	}
	for _, item := range src {
		inst.add(item)
	}
	return nil
}

func (inst *myPakcageUninstaller) add(o any) {
	if o == nil {
		return
	}
	inst.objects = append(inst.objects, o)
}

func (inst *myPakcageUninstaller) addFile(o *entity.InstalledFile) {
	if o == nil {
		return
	}
	inst.files = append(inst.files, o)
}

func (inst *myPakcageUninstaller) tryRemoveDir(f *entity.InstalledFile) error {
	path := inst.FileSystemSer.Path(f.Path)
	if !path.IsDirectory() {
		return nil // skip
	}
	err := path.Delete()
	if err != nil {
		vlog.Warn(err)
	}
	return nil
}

func (inst *myPakcageUninstaller) tryRemoveFile(f *entity.InstalledFile) error {

	path := inst.FileSystemSer.Path(f.Path)
	if !path.IsFile() {
		return nil // skip
	}

	wantSum := f.SHA256SUM
	wantSize := f.Size
	haveSize := path.GetInfo().Length()
	haveSum, err := utils.ComputeFileSHA256sumAFS(path)
	if err != nil {
		return err
	}

	// check size
	if wantSize != haveSize {
		return nil // skip
	}

	// check sum
	if wantSum != haveSum {
		return nil // skip
	}

	return path.Delete()
}

func (inst *myPakcageUninstaller) removeFiles() error {
	list := inst.files

	sa := &utils.SortAdapter{Quietly: true}
	sa.OnLen = func() int {
		return len(list)
	}
	sa.OnLess = func(i1, i2 int) bool {
		p1 := list[i1].Path
		p2 := list[i2].Path
		return strings.Compare(p1, p2) > 0
	}
	sa.OnSwap = func(i1, i2 int) {
		o1 := list[i1]
		o2 := list[i2]
		list[i1] = o2
		list[i2] = o1
	}
	sa.Sort()

	var err error = nil
	for _, f := range list {
		if f.IsFile {
			err = inst.tryRemoveFile(f)
		} else if f.IsDir {
			err = inst.tryRemoveDir(f)
		} else {
			continue
		}
		if err != nil {
			vlog.Warn(err)
		}
	}
	return nil
}

func (inst *myPakcageUninstaller) removeObjects() error {

	db := inst.db
	all := inst.objects
	elist := utils.ErrorList{}

	if db == nil {
		return fmt.Errorf("no gorm.DB instance")
	}

	for _, item := range all {
		res := db.Delete(item)
		elist.Append(res.Error)
	}

	return elist.First()
}

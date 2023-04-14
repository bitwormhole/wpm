package plugins

import (
	"context"
	"fmt"

	"github.com/bitwormhole/wpm/server/data/dbagent"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/utils"
	"gorm.io/gorm"
)

type myPakcageUninstaller struct {
	context      context.Context
	agent        dbagent.GormDBAgent
	installation dxo.InstallationID

	// pack    *dto.SoftwarePackage

	// HTTPClient   service.HTTPClientService
	// HTTPClientEx service.HTTPClientExService

	// ContentTypeSer    service.ContentTypeService
	// MediaSer          service.MediaService
	// IntentTemplateSer service.IntentTemplateService
	// ExecutableSer     service.ExecutableService

	db      *gorm.DB
	objects []any
}

func (inst *myPakcageUninstaller) Uninstall() error {

	steps := make([]func() error, 0)

	steps = append(steps, inst.prepare)

	steps = append(steps, inst.findContentTypes)
	steps = append(steps, inst.findExecutables)
	steps = append(steps, inst.findIntentTemplates)
	steps = append(steps, inst.findMediae)

	steps = append(steps, inst.remove)

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

func (inst *myPakcageUninstaller) remove() error {

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

package setup

import (
	"context"
	"fmt"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ImpSetupService ...
type ImpSetupService struct {
	markup.Component `id:"SetupService"`

	AppDataService           service.AppDataService           `inject:"#AppDataService"`
	FileSystemService        service.FileSystemService        `inject:"#FileSystemService"`
	ExecutableService        service.ExecutableService        `inject:"#ExecutableService"`
	ExecutableImportService  service.ExecutableImportService  `inject:"#ExecutableImportService"`
	IntentTemplateService    service.IntentTemplateService    `inject:"#IntentTemplateService"`
	MediaService             service.MediaService             `inject:"#MediaService"`
	ProjectTypeImportService service.ProjectTypeImportService `inject:"#ProjectTypeImportService"`
	SettingService           service.SettingService           `inject:"#SettingService"`

	cachedRegs []*service.SetupRegistration
}

func (inst *ImpSetupService) _Impl() service.SetupService {
	return inst
}

// IsSetupReqiured ...
func (inst *ImpSetupService) IsSetupReqiured(ctx context.Context) (bool, error) {
	sett, err := inst.SettingService.GetSettings()
	if err != nil {
		return false, err
	}
	done := sett.SetupDone
	return !done, nil
}

// Apply ...
func (inst *ImpSetupService) Apply(ctx context.Context, items []*dto.Setup) error {
	for _, item := range items {
		err := inst.applyItem(ctx, item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *ImpSetupService) applyItem(ctx context.Context, item *dto.Setup) error {
	reg, err := inst.findRegistration(item)
	if err != nil {
		return err
	}
	err = reg.Handler(ctx, item)
	if err == nil {
		item.State = dto.SetupStateSuccess
		item.Error = ""
	} else {
		item.State = dto.SetupStateError
		item.Error = err.Error()
	}
	return nil
}

func (inst *ImpSetupService) findRegistration(item *dto.Setup) (*service.SetupRegistration, error) {
	id := item.ID
	name := item.Name
	all := inst.listSetupRegs()
	for _, reg := range all {
		if reg.Name == name && reg.ID == id {
			return reg, nil
		}
	}
	return nil, fmt.Errorf("no setup item with id [%v] and name [%v]", id, name)
}

// ListAll ...
func (inst *ImpSetupService) ListAll(ctx context.Context) ([]*dto.Setup, error) {
	src := inst.listSetupRegs()
	dst := make([]*dto.Setup, 0)
	for _, reg := range src {
		o := reg.Prototype
		if o == nil {
			continue
		}
		dst = append(dst, o)
	}
	return dst, nil
}

func (inst *ImpSetupService) listSetupRegs() []*service.SetupRegistration {
	all := inst.cachedRegs
	if all != nil {
		return all
	}
	ctx := inst.makeContext()
	reg := &mySetupReg{context: ctx}
	all = reg.listAll()
	inst.cachedRegs = all
	return all
}

func (inst *ImpSetupService) makeContext() *Context {
	return &Context{
		AppDataService:           inst.AppDataService,
		FileSystemService:        inst.FileSystemService,
		ExecutableImportService:  inst.ExecutableImportService,
		IntentTemplateService:    inst.IntentTemplateService,
		MediaService:             inst.MediaService,
		ProjectTypeImportService: inst.ProjectTypeImportService,
		SettingService:           inst.SettingService,
	}
}

// SkipAll ...
func (inst *ImpSetupService) SkipAll(ctx context.Context) error {
	ser := inst.SettingService
	sett, err := ser.GetSettings()
	if err != nil {
		return err
	}
	sett.SetupDone = true
	return ser.PutSettings(sett)
}

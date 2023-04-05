package setup

import (
	"context"
	"fmt"

	"github.com/bitwormhole/wpm/server/web/dto"
)

type mySetupHanlders struct {
	context *Context
}

func (inst *mySetupHanlders) example(c context.Context, req *dto.Setup) error {
	return fmt.Errorf("no impl")
}

func (inst *mySetupHanlders) doImportPresetProjectTypes(c context.Context, req *dto.Setup) error {
	ser := inst.context.ProjectTypeImportService
	return ser.ImportTypesFromPreset(c)
}

func (inst *mySetupHanlders) doImportPresetMediaFiles(c context.Context, req *dto.Setup) error {
	ser := inst.context.MediaService
	return ser.ImportPresets(c)
}

func (inst *mySetupHanlders) doImportPresetIntentTemplates(c context.Context, req *dto.Setup) error {
	ser := inst.context.IntentTemplateService
	return ser.ImportPreset(c)
}

func (inst *mySetupHanlders) doImportCommonExecutables(c context.Context, req *dto.Setup) error {
	ser := inst.context.ExecutableImportService
	return ser.ImportPresets(c)
}

func (inst *mySetupHanlders) doImportPresetSettings(c context.Context, req *dto.Setup) error {
	ser := inst.context.SettingService
	sett, err := ser.GetSettings()
	if err != nil {
		return err
	}
	sett.SetupDone = true
	return ser.PutSettings(sett)
}

func (inst *mySetupHanlders) doInitMainRepository(c context.Context, req *dto.Setup) error {
	h := &myInitMainRepositoryHandler{
		context: inst.context,
	}
	return h.Run()
}

func (inst *mySetupHanlders) doDumpOlderData(c context.Context, req *dto.Setup) error {
	return fmt.Errorf("todo: no impl")
}

func (inst *mySetupHanlders) doImportOlderData(c context.Context, req *dto.Setup) error {
	return fmt.Errorf("todo: no impl")
}

package intenttemplates

import (
	"context"
	"fmt"

	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/web/dto"
)

type myPresetTemplatesImport struct {
	parent *IntentTemplateServiceImpl
}

func (inst *myPresetTemplatesImport) Import(c context.Context) error {

	list, err := inst.loadTemplates()
	if err != nil {
		return err
	}

	for _, item := range list {
		_, err = inst.parent.Insert(c, item)
		if err != nil {
			vlog.Warn(err)
		}
	}

	return nil
}

func (inst *myPresetTemplatesImport) loadTemplates() ([]*dto.IntentTemplate, error) {

	ser := inst.parent.PresetService
	all, err := ser.GetPresets()
	if err != nil {
		return nil, err
	}

	list := all.IntentTemplates
	if len(list) < 1 {
		return nil, fmt.Errorf("no preset Intent-Template")
	}

	return list, nil
}

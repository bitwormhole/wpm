package intenttemplates

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/backup"
	"github.com/bitwormhole/wpm/server/data/entity"
)

type myPresetTemplatesImport struct {
	parent *IntentTemplateServiceImpl
}

func (inst *myPresetTemplatesImport) Import(c context.Context) error {

	list1, err := inst.loadTemplates("preset-intent-templates.json")
	if err != nil {
		return err
	}

	list2, err := inst.parent.entity2dtoList(list1)
	if err != nil {
		return err
	}

	for _, item := range list2 {
		_, err = inst.parent.Insert(c, item)
		if err != nil {
			vlog.Warn(err)
		}
	}

	return nil
}

func (inst *myPresetTemplatesImport) loadTemplates(name string) ([]*entity.IntentTemplate, error) {

	url := "res:///" + name
	data, err := inst.parent.AC.GetResources().GetBinary(url)
	if err != nil {
		return nil, err
	}

	vo := &backup.VO{}
	err = json.Unmarshal(data, vo)
	if err != nil {
		return nil, err
	}

	list := vo.IntentTemplateTable
	if list == nil {
		return nil, fmt.Errorf("the 'IntentTemplateTable' is nil")
	}

	return list, nil
}

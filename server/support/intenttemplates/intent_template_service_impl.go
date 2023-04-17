package intenttemplates

import (
	"context"
	"errors"
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/components/intents"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// IntentTemplateServiceImpl ...
type IntentTemplateServiceImpl struct {
	markup.Component `id:"IntentTemplateService"`

	AC                  application.Context               `inject:"context"`
	IntentTempDAO       dao.IntentTemplateDAO             `inject:"#IntentTemplateDAO"`
	IntentFilterManager intents.FilterManager             `inject:"#wpm-intent-filter-manager"`
	PresetService       service.PresetService             `inject:"#PresetService"`
	TemplateCache       service.IntentTemplateEntityCache `inject:"#IntentTemplateEntityCache"`
}

func (inst *IntentTemplateServiceImpl) _Impl() service.IntentTemplateService {
	return inst
}

func (inst *IntentTemplateServiceImpl) normalizeText(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	return text
}

func (inst *IntentTemplateServiceImpl) dto2entity(o1 *dto.IntentTemplate) (*entity.IntentTemplate, error) {

	action := o1.ActionRequest
	selbuilder := dxo.IntentTemplateSelectorBuilder{
		Method: action.Method,
		Target: action.Target,
		Type:   action.Type,
		With:   action.With.String(),
	}

	o2 := &entity.IntentTemplate{}
	o2.ID = o1.ID
	o2.UUID = o1.UUID
	o2.Installation = o1.Installation

	o2.Name = o1.Name
	o2.Title = o1.Title
	o2.Group = o1.Group
	o2.IconURL = o1.IconURL
	o2.Description = o1.Description

	o2.Executable = action.With
	o2.Method = inst.normalizeText(action.Method)
	o2.Target = inst.normalizeText(action.Target)
	o2.ContentType = inst.normalizeText(action.Type)
	o2.Selector = selbuilder.Create()

	if o2.Target != "file" && o2.Target != "project" {
		o2.ContentType = "*"
	}

	o2.Arguments = o1.Arguments
	o2.Command = o1.Command
	o2.Env = o1.Env
	o2.WD = o1.WD

	// todo ...
	return o2, nil
}

func (inst *IntentTemplateServiceImpl) entity2dto(o1 *entity.IntentTemplate) (*dto.IntentTemplate, error) {

	o2 := &dto.IntentTemplate{}
	o2.ID = o1.ID
	o2.UUID = o1.UUID
	o2.Installation = o1.Installation
	o2.CreatedAt = util.NewTime(o1.CreatedAt)
	o2.UpdatedAt = util.NewTime(o1.UpdatedAt)

	o2.Name = o1.Name
	o2.Title = o1.Title
	o2.Group = o1.Group
	o2.Description = o1.Description
	o2.IconURL = o1.IconURL

	o2.With = o1.Executable
	o2.Method = inst.normalizeText(o1.Method)
	o2.Target = inst.normalizeText(o1.Target)
	o2.Type = inst.normalizeText(o1.ContentType)
	o2.Selector = o1.Selector

	o2.Arguments = o1.Arguments
	o2.Command = o1.Command
	o2.Env = o1.Env
	o2.WD = o1.WD

	// todo ...
	return o2, nil
}

func (inst *IntentTemplateServiceImpl) entity2dtoList(src []*entity.IntentTemplate) ([]*dto.IntentTemplate, error) {
	dst := make([]*dto.IntentTemplate, 0)
	for _, item1 := range src {
		item2, err := inst.entity2dto(item1)
		if err != nil {
			return nil, err
		}
		dst = append(dst, item2)
	}
	return dst, nil
}

// ListAll ...
func (inst *IntentTemplateServiceImpl) ListAll(ctx context.Context) ([]*dto.IntentTemplate, error) {
	src, err := inst.IntentTempDAO.ListAll()
	if err != nil {
		return nil, err
	}
	return inst.entity2dtoList(src)
}

// ListBySelector ...
func (inst *IntentTemplateServiceImpl) ListBySelector(ctx context.Context, sel *dto.IntentTemplate) ([]*dto.IntentTemplate, error) {

	want, err := inst.dto2entity(sel)
	if err != nil {
		return nil, err
	}

	all, err := inst.TemplateCache.ListTemplates()
	if err != nil {
		return nil, err
	}

	dst := make([]*entity.IntentTemplate, 0)
	for _, item := range all {
		if inst.isTemplateMatch(item, want) {
			dst = append(dst, item)
		}
	}

	return inst.entity2dtoList(dst)
}

func (inst *IntentTemplateServiceImpl) isTemplateMatch(have, want *entity.IntentTemplate) bool {
	b1 := inst.isTextMatch(have.Method, want.Method)
	b2 := inst.isTextMatch(have.Target, want.Target)
	b3 := inst.isTextMatch(have.ContentType, want.ContentType)
	b4 := inst.isTextMatch(have.Executable.String(), want.Executable.String())
	return (b1 && b2 && b3 && b4)
}

func (inst *IntentTemplateServiceImpl) isTextMatch(have, want string) bool {
	if have == "*" || want == "*" {
		return true
	}
	return have == want
}

// Find ...
func (inst *IntentTemplateServiceImpl) Find(ctx context.Context, id dxo.IntentTemplateID) (*dto.IntentTemplate, error) {
	return nil, errors.New("no impl")
}

// Insert ...
func (inst *IntentTemplateServiceImpl) Insert(ctx context.Context, o1 *dto.IntentTemplate) (*dto.IntentTemplate, error) {
	o2, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}
	o3, err := inst.IntentTempDAO.Insert(o2)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o3)
}

// Update ...
func (inst *IntentTemplateServiceImpl) Update(ctx context.Context, id dxo.IntentTemplateID, o1 *dto.IntentTemplate) (*dto.IntentTemplate, error) {
	o2, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}
	o3, err := inst.IntentTempDAO.Update(id, o2)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o3)
}

// Remove ...
func (inst *IntentTemplateServiceImpl) Remove(ctx context.Context, id dxo.IntentTemplateID) error {
	return inst.IntentTempDAO.Remove(id)
}

// ImportPreset ...
func (inst *IntentTemplateServiceImpl) ImportPreset(ctx context.Context) error {
	im := &myPresetTemplatesImport{parent: inst}
	return im.Import(ctx)
}

// ListMacroProperties ...
func (inst *IntentTemplateServiceImpl) ListMacroProperties(ctx context.Context) (map[string]string, error) {

	const path = "res:///presets/common/intent-macro-list.properties"
	text, err := inst.AC.GetResources().GetText(path)
	if err != nil {
		return nil, err
	}

	props, err := collection.ParseProperties(text, nil)
	if err != nil {
		return nil, err
	}

	src := intents.ListMacroNames()
	dst := make(map[string]string)
	for _, name := range src {
		value := props.GetProperty(name, name)
		dst[name] = value
	}
	return dst, nil
}

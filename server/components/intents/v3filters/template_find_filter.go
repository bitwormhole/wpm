package v3filters

import (
	"context"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/components/intents"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// FindTemplateFilter ...
type FindTemplateFilter struct {
	markup.Component `class:"wpm-intent-filter"`

	IntentTemplateService service.IntentTemplateService `inject:"#IntentTemplateService"`
}

func (inst *FindTemplateFilter) _Impl() (intents.FilterRegistry, intents.Filter) {
	return inst, inst
}

// GetFilterRegistrationList ...
func (inst *FindTemplateFilter) GetFilterRegistrationList() []*intents.FilterRegistration {
	fr := &intents.FilterRegistration{
		Enabled: true,
		Order:   orderFindTemplate,
		Filter:  inst,
	}
	return []*intents.FilterRegistration{fr}
}

// Handle ...
func (inst *FindTemplateFilter) Handle(c context.Context, i *dto.Intent, next intents.FilterChain) error {

	templist := i.Templates
	if len(templist) > 0 {
		return next.Handle(c, i)
	}

	sel := &dto.IntentTemplate{}
	sel.ActionRequest = i.Action
	list, err := inst.IntentTemplateService.ListBySelector(c, sel)
	if err != nil {
		return err
	}

	i.Templates = list
	return next.Handle(c, i)
}

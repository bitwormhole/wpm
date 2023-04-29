package v3filters

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/components/intents"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// CheckTemplateFilter ...
type CheckTemplateFilter struct {
	markup.Component `class:"wpm-intent-filter"`
}

func (inst *CheckTemplateFilter) _Impl() (intents.FilterRegistry, intents.Filter) {
	return inst, inst
}

// GetFilterRegistrationList ...
func (inst *CheckTemplateFilter) GetFilterRegistrationList() []*intents.FilterRegistration {
	fr := &intents.FilterRegistration{
		Enabled: true,
		Order:   orderCheckTemplate,
		Filter:  inst,
	}
	return []*intents.FilterRegistration{fr}
}

// Handle ...
func (inst *CheckTemplateFilter) Handle(c context.Context, i *dto.Intent, next intents.FilterChain) error {

	temp := i.Template
	want := i.Want
	have := i.Have
	count := 0

	if temp == nil {
		templist := make([]*dto.IntentTemplate, 0)
		if want != nil {
			templist = append(templist, want.Templates...)
		}
		if have != nil {
			templist = append(templist, have.Templates...)
		}
		for _, item := range templist {
			if item == nil {
				continue
			}
			count++
			temp = item
		}
	}

	if count > 1 {
		i.Status = http.StatusContinue
		i.Message = "multi-intent-templates"
		return nil
	}

	if temp == nil {
		return fmt.Errorf("no template for intent, selector = %v", i.Selector)
	}

	i.Template = temp
	return next.Handle(c, i)
}

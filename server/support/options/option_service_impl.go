package options

import (
	"fmt"
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// ImpOptionService ...
type ImpOptionService struct {
	markup.Component `id:"OptionService"`

	AC application.Context `inject:"context"`
}

func (inst *ImpOptionService) _Impl() service.OptionService {
	return inst
}

// GetOptions ...
func (inst *ImpOptionService) GetOptions(o *vo.Option) error {
	if o == nil {
		return fmt.Errorf("param:o is nil")
	}
	o.Options = inst.listOptions()
	return nil
}

func (inst *ImpOptionService) listOptions() map[string]bool {
	const prefix = "wpm.options."
	props := inst.AC.GetProperties()
	g := props.Getter()
	src := props.Export(nil)
	dst := make(map[string]bool)
	for k := range src {
		if strings.HasPrefix(k, prefix) {
			dst[k] = g.GetBool(k, false)
		}
	}
	return dst
}

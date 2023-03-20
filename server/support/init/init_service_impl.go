package init

import (
	"context"
	"fmt"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// ImpInitService ...
type ImpInitService struct {
	markup.Component `id:"InitService"`

	ProjectTypeService service.ProjectTypeService `inject:"#ProjectTypeService"`
	ExecutableService  service.ExecutableService  `inject:"#ExecutableService"`
	CheckUpdateService service.CheckUpdateService `inject:"#CheckUpdateService"`
	SetupService       service.SetupService       `inject:"#SetupService"`
}

func (inst *ImpInitService) _Impl() service.InitService {
	return inst
}

// InitGet ...
func (inst *ImpInitService) InitGet(ctx context.Context) (*vo.Init, error) {
	h := &myInitGetHandler{
		ProjectTypeSer: inst.ProjectTypeService,
		ExecutableSer:  inst.ExecutableService,
		CheckUpdateSer: inst.CheckUpdateService,
		SetupSer:       inst.SetupService,
	}
	return h.handle(ctx)
}

// InitSet ...
func (inst *ImpInitService) InitSet(ctx context.Context, o *vo.Init) (*vo.Init, error) {
	return nil, fmt.Errorf("no impl")
}

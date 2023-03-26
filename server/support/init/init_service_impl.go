package init

import (
	"context"
	"fmt"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// ImpInitService ...
type ImpInitService struct {
	markup.Component `id:"InitService" class:"life"`

	AboutService       service.AboutService       `inject:"#AboutService"`
	ProjectTypeService service.ProjectTypeService `inject:"#ProjectTypeService"`
	ExecutableService  service.ExecutableService  `inject:"#ExecutableService"`
	CheckUpdateService service.CheckUpdateService `inject:"#CheckUpdateService"`
	SetupService       service.SetupService       `inject:"#SetupService"`
}

func (inst *ImpInitService) _Impl() (service.InitService, application.LifeRegistry) {
	return inst, inst
}

// GetLifeRegistration ...
func (inst *ImpInitService) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnInit:   inst.init,
		Priority: 1000,
	}
}

func (inst *ImpInitService) init() error {
	release := inst.AboutService.IsRelease()
	if release {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	return nil
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

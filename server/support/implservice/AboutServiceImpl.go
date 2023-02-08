package implservice

import (
	"context"
	"fmt"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// AboutServiceImpl ...
type AboutServiceImpl struct {
	markup.Component `id:"AboutService"`

	Profile string `inject:"${application.profiles.active}"`

	Name      string `inject:"${application.about.name}"`
	Title     string `inject:"${application.about.title}"`
	Copyright string `inject:"${application.about.copyright}"`

	PlatformService service.PlatformService `inject:"#PlatformService"`
	ProfileService  service.ProfileService  `inject:"#ProfileService"`
}

func (inst *AboutServiceImpl) _Impl() service.AboutService {
	return inst
}

// GetInfo ...
func (inst *AboutServiceImpl) GetInfo(ctx context.Context) (*vo.About, error) {

	o := &vo.About{}

	o.Name = inst.Name
	o.Title = inst.Title
	o.Copyright = inst.Copyright

	inst.loadPlatform(ctx, o)
	inst.loadProfile(ctx, o)
	inst.loadAppModule(ctx, o)

	return o, nil
}

func (inst *AboutServiceImpl) loadPlatform(ctx context.Context, view *vo.About) error {
	p, err := inst.PlatformService.GetPlatform()
	if err != nil {
		return err
	}
	view.Arch = p.Arch
	view.OS = p.OS
	return nil
}

func (inst *AboutServiceImpl) loadProfile(ctx context.Context, view *vo.About) error {
	view.Profile = inst.Profile
	p, err := inst.ProfileService.GetProfile()
	if err != nil {
		return err
	}
	view.Home = p.Home
	view.User = p.User
	return nil
}

func (inst *AboutServiceImpl) loadAppModule(ctx context.Context, view *vo.About) error {
	m := service.GetAppModule()
	if m == nil {
		return fmt.Errorf("no module info")
	}
	view.Module = m.GetName()
	view.Version = m.GetVersion()
	view.Revision = m.GetRevision()
	return nil
}

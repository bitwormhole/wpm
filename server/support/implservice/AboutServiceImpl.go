package implservice

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// AboutServiceImpl ...
type AboutServiceImpl struct {
	markup.Component `id:"AboutService"`

	Profile string `inject:"${application.profiles.active}"`

	Name       string `inject:"${application.about.name}"`
	Title      string `inject:"${application.about.title}"`
	Copyright  string `inject:"${application.about.copyright}"`
	ServerPort int    `inject:"${server.port}"`

	EnableDebug bool `inject:"${wpm.debug.enabled}"`

	PlatformService   service.PlatformService   `inject:"#PlatformService"`
	ProfileService    service.ProfileService    `inject:"#ProfileService"`
	AppRuntimeService service.AppRuntimeService `inject:"#AppRuntimeService"`
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
	inst.loadWebInfo(ctx, o)
	inst.loadAppRuntimeInfo(ctx, o)

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
	loader := &myAboutModuleInfoLoader{}
	return loader.load(view)
}

func (inst *AboutServiceImpl) loadWebInfo(ctx context.Context, view *vo.About) error {
	port := strconv.Itoa(inst.ServerPort)
	view.WebURL = "http://localhost:" + port
	return nil
}

func (inst *AboutServiceImpl) loadAppRuntimeInfo(ctx context.Context, view *vo.About) error {
	sum, err := inst.AppRuntimeService.GetAppHash()
	if err != nil {
		return err
	}
	view.SHA256SUM = sum
	return nil
}

// IsDebug ...
func (inst *AboutServiceImpl) IsDebug() bool {
	return inst.EnableDebug
}

// IsRelease ...
func (inst *AboutServiceImpl) IsRelease() bool {
	return !inst.EnableDebug
}

////////////////////////////////////////////////////////////////////////////////

type myAboutModuleInfoLoader struct {
}

func (inst *myAboutModuleInfoLoader) load(view *vo.About) error {

	mm1, err := inst.getMainModule()
	if err != nil {
		return err
	}
	mm2 := inst.makeDTO(mm1)

	deplist1 := inst.getDepList(mm1)
	deplist2 := inst.makeDTOs(deplist1)

	view.MainModule = *mm2
	view.Module = mm2.Name
	view.Version = mm2.Version
	view.Revision = mm2.Revision
	view.Modules = deplist2
	return nil
}

func (inst *myAboutModuleInfoLoader) makeDTO(src application.Module) *dto.Module {
	dst := &dto.Module{}
	if src == nil {
		return dst
	}

	name := src.GetName()
	// hName := utils.ComputeSHA256sum([]byte(name))

	dst.Name = name
	dst.Version = src.GetVersion()
	dst.Revision = src.GetRevision()
	// dst.HexName = hName.String()

	return dst
}

func (inst *myAboutModuleInfoLoader) makeDTOs(src []application.Module) []*dto.Module {
	dst := make([]*dto.Module, 0)
	for _, m1 := range src {
		m2 := inst.makeDTO(m1)
		dst = append(dst, m2)
	}
	return dst
}

func (inst *myAboutModuleInfoLoader) getMainModule() (application.Module, error) {
	m := service.GetAppModule()
	if m == nil {
		return nil, fmt.Errorf("no module info")
	}
	return m, nil
}

func (inst *myAboutModuleInfoLoader) getDepList(mm application.Module) []application.Module {
	const limit = 64
	mods := make([]application.Module, 0)
	table := make(map[string]application.Module)
	err := inst.walkDepTree(mm, table, limit)
	if err != nil {
		return mods
	}
	keys := make([]string, 0)
	for key := range table {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		m := table[key]
		mods = append(mods, m)
	}
	return mods
}

func (inst *myAboutModuleInfoLoader) walkDepTree(m application.Module, dst map[string]application.Module, limit int) error {

	if limit < 0 {
		return fmt.Errorf("the dep tree is too deep")
	}

	if m == nil || dst == nil {
		return fmt.Errorf("param is nil")
	}

	// for self info
	name := m.GetName()
	ver := m.GetVersion()
	key := name + "@" + ver
	old := dst[key]
	if old == nil {
		dst[key] = m
	} else {
		return nil
	}

	// for children
	deplist := m.GetDependencies()
	for _, d := range deplist {
		err := inst.walkDepTree(d, dst, limit-1)
		if err != nil {
			return err
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

package implservice

import (
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// AppIconServiceImpl ...
type AppIconServiceImpl struct {
	markup.Component `id:"AppIconService" class:"life"`

	PropsName string              `inject:"${wpm.exe-icons.properties}"`
	Context   application.Context `inject:"context"`

	table map[string]string
}

func (inst *AppIconServiceImpl) _Impl() (service.AppIconService, application.LifeRegistry) {
	return inst, inst
}

// GetLifeRegistration ...
func (inst *AppIconServiceImpl) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnInit: inst.init,
	}
}

func (inst *AppIconServiceImpl) init() error {

	dst := make(map[string]string)
	file := inst.PropsName
	text, err := inst.Context.GetResources().GetText(file)
	if err != nil {
		return err
	}

	src, err := collection.ParseProperties(text, nil)
	if err != nil {
		return err
	}

	dst = src.Export(dst)
	inst.table = dst
	return nil
}

// FillWithIconURL ...
func (inst *AppIconServiceImpl) FillWithIconURL(o1 *dto.Executable) error {
	name := o1.Name
	url := inst.findIconURL(name)
	o1.IconURL = url
	return nil
}

func (inst *AppIconServiceImpl) normalizeName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	i := strings.LastIndex(name, ".")
	if i > 0 {
		name = name[0:i]
	}
	return name
}

func (inst *AppIconServiceImpl) findIconURL(name string) string {

	name = inst.normalizeName(name)
	key := "icon." + name + ".url"
	url := inst.table[key]

	if url == "" {
		key = "icon.default.url"
		url = inst.table[key]
	}

	if url == "" {
		url = "/image/exe-icons/default.png"
	}

	return url
}

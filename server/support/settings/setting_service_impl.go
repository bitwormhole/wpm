package settings

import (
	"time"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/settings"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// SettingServiceImpl ...
type SettingServiceImpl struct {
	markup.Component `id:"SettingService" class:"life" `

	SettingDAO dao.SettingDAO `inject:"#SettingDAO"`
}

func (inst *SettingServiceImpl) _Impl() (service.SettingService, application.LifeRegistry) {
	return inst, inst
}

// GetLifeRegistration ...
func (inst *SettingServiceImpl) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnStart: inst.onStart,
	}
}

func (inst *SettingServiceImpl) onStart() error {

	now := time.Now()
	inst.PutString("a", "1")
	inst.PutString("b", "2")
	inst.PutString("t", now.String())

	return nil
}

// ListAll ...
func (inst *SettingServiceImpl) ListAll() map[string]string {
	dst := make(map[string]string)
	src, err := inst.SettingDAO.ListAll()
	if err != nil {
		return dst
	}
	for _, item := range src {
		name := item.Name
		value := item.Value
		dst[name] = value
	}
	return dst
}

// PutString ...
func (inst *SettingServiceImpl) PutString(name, value string) {
	o := &entity.Setting{
		Name:  name,
		Value: value,
	}
	inst.SettingDAO.Put(name, o)
}

// GetString ...
func (inst *SettingServiceImpl) GetString(name, defaultValue string) string {
	item := inst.SettingDAO.Get(name, nil)
	if item != nil {
		return item.Value
	}
	return ""
}

// PutSettings ...
func (inst *SettingServiceImpl) PutSettings(o *dto.Settings) error {
	src := inst.dto2map(o)
	table := make(map[string]*entity.Setting)
	dst, err := inst.SettingDAO.ListAll()
	if err == nil {
		for _, item := range dst {
			table[item.Name] = item
		}
	}
	for k, v := range src {
		item := table[k]
		if item == nil {
			// do insert
			item = &entity.Setting{
				Name:  k,
				Value: v,
			}
			inst.SettingDAO.Insert(item)
		} else {
			// do update
			item.Value = v
			inst.SettingDAO.Update(item.ID, item)
		}
	}
	return nil
}

// GetSettings ...
func (inst *SettingServiceImpl) GetSettings() (*dto.Settings, error) {
	all := inst.ListAll()
	o := inst.map2dto(all)
	return o, nil
}

func (inst *SettingServiceImpl) dto2map(src *dto.Settings) map[string]string {
	dst := make(map[string]string)
	if src == nil {
		return dst
	}

	p := collection.CreateProperties()
	s := p.Setter()

	s.SetInt(settings.IgnorePackageRevision, src.IgnorePackageRevision)
	s.SetString(settings.IgnorePackageVersion, src.IgnorePackageVersion)
	s.SetString(settings.IgnorePackageSum, src.IgnorePackageSum.String())
	s.SetBool(settings.SetupDone, src.SetupDone)

	return p.Export(dst)
}

func (inst *SettingServiceImpl) map2dto(src map[string]string) *dto.Settings {
	dst := &dto.Settings{}
	if src == nil {
		return dst
	}

	p := collection.CreateProperties()
	p.Import(src)
	g := p.Getter()

	dst.IgnorePackageRevision = g.GetInt(settings.IgnorePackageRevision, 0)
	dst.IgnorePackageVersion = g.GetString(settings.IgnorePackageVersion, "")
	dst.IgnorePackageSum = util.Hex(g.GetString(settings.IgnorePackageSum, ""))
	dst.SetupDone = g.GetBool(settings.SetupDone, false)

	return dst
}

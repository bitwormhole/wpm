package checkupdate

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// TheCheckUpdateServiceImpl ...
type TheCheckUpdateServiceImpl struct {
	markup.Component `id:"CheckUpdateService"`

	PackagesURL    string                 `inject:"${wpm.check-update.url}"`
	AboutService   service.AboutService   `inject:"#AboutService"`
	SettingService service.SettingService `inject:"#SettingService"`

	cached *vo.AboutCheckUpdate
}

func (inst *TheCheckUpdateServiceImpl) _Impl() service.CheckUpdateService {
	return inst
}

// Check ...
func (inst *TheCheckUpdateServiceImpl) Check(ctx context.Context, o *vo.AboutCheckUpdate) error {

	if o.Auto && inst.cached != nil {
		o.HasIgnored = false
		o.HasNewVersion = false
		return nil
	}
	inst.cached = o

	err := inst.checkCurrent(ctx, o)
	if err != nil {
		return err
	}

	err = inst.checkIgnore(ctx, o)
	if err != nil {
		return err
	}

	err = inst.checkLatest(ctx, o)
	if err != nil {
		return err
	}

	o.HasIgnored = inst.computeHasIgnored(o)
	o.HasNewVersion = inst.computeHasNewVersion(o)

	return nil
}

func (inst *TheCheckUpdateServiceImpl) computeHasNewVersion(o *vo.AboutCheckUpdate) bool {
	cur := o.Current
	lat := o.Latest
	if cur == nil || lat == nil {
		return false
	}
	b1 := cur.ModuleName == lat.ModuleName
	b2 := cur.OS == lat.OS
	b3 := cur.Arch == lat.Arch
	b4 := cur.Revision < lat.Revision
	return (b1 && b2 && b3 && b4)
}

func (inst *TheCheckUpdateServiceImpl) computeHasIgnored(o *vo.AboutCheckUpdate) bool {
	ign := o.Ignore
	lat := o.Latest
	if ign == nil || lat == nil {
		return false
	}
	b1 := ign.Revision == lat.Revision
	b2 := ign.Version == lat.Version
	b3 := ign.SHA256SUM == lat.SHA256SUM
	return (b1 && b2 && b3)
}

// Ignore ...
func (inst *TheCheckUpdateServiceImpl) Ignore(ctx context.Context, o *vo.AboutCheckUpdate) error {

	ig := o.Ignore
	ss := inst.SettingService
	sett, err := ss.GetSettings()
	if err != nil {
		return err
	}

	sett.IgnorePackageRevision = ig.Revision
	sett.IgnorePackageSum = ig.SHA256SUM
	sett.IgnorePackageVersion = ig.Version

	return ss.PutSettings(sett)
}

func (inst *TheCheckUpdateServiceImpl) checkCurrent(ctx context.Context, o *vo.AboutCheckUpdate) error {

	info, err := inst.AboutService.GetInfo(ctx)
	if err != nil {
		return err
	}

	sp := &dto.SoftwarePackage{}
	sp.ModuleName = info.MainModule.Name
	sp.Version = info.MainModule.Version
	sp.Revision = info.MainModule.Revision
	sp.OS = info.OS
	sp.Arch = info.Arch

	o.Current = sp
	return nil
}

func (inst *TheCheckUpdateServiceImpl) checkLatest(ctx context.Context, o *vo.AboutCheckUpdate) error {

	latest, err := inst.fetchLatest(ctx)
	if err != nil {
		return err
	}

	current := o.Current
	sp := current
	all := latest.Packages
	for _, have := range all {
		if inst.isMatch(current, have) {
			if current.Revision < have.Revision {
				sp = have
				break
			}
		}
	}

	o.Latest = sp
	return nil
}

func (inst *TheCheckUpdateServiceImpl) checkIgnore(ctx context.Context, o *vo.AboutCheckUpdate) error {

	sett, err := inst.SettingService.GetSettings()
	if sett == nil || err != nil {
		return nil
	}

	curr := o.Current

	sp := &dto.SoftwarePackage{}
	sp.Version = sett.IgnorePackageVersion
	sp.Revision = sett.IgnorePackageRevision
	sp.SHA256SUM = sett.IgnorePackageSum
	sp.Arch = curr.Arch
	sp.OS = curr.OS

	o.Ignore = sp
	return nil
}

func (inst *TheCheckUpdateServiceImpl) fetchLatest(ctx context.Context) (*vo.SoftwarePackage, error) {

	url := inst.PackagesURL
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body := resp.Body
	defer func() {
		if body != nil {
			body.Close()
		}
	}()

	code := resp.StatusCode
	if code != http.StatusOK {
		return nil, fmt.Errorf("HTTP " + resp.Status)
	}

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	sp := &vo.SoftwarePackage{}
	err = json.Unmarshal(data, sp)
	if err != nil {
		return nil, err
	}

	return sp, nil
}

func (inst *TheCheckUpdateServiceImpl) isMatch(p1, p2 *dto.SoftwarePackage) bool {

	if p1 == nil || p2 == nil {
		return false
	}

	if p1.OS != p2.OS {
		return false
	}

	if p1.Arch != p2.Arch {
		return false
	}

	if p1.ModuleName != p2.ModuleName {
		return false
	}

	return true
}

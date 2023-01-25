package implservice

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ProfileServiceImpl ...
type ProfileServiceImpl struct {
	markup.Component `id:"ProfileService"`

	PlatformService service.PlatformService `inject:"#PlatformService"`
}

func (inst *ProfileServiceImpl) _Impl() service.ProfileService {
	return inst
}

// GetProfile ...
func (inst *ProfileServiceImpl) GetProfile() (*dto.Profile, error) {
	provider, err := inst.PlatformService.GetProvider()
	if err != nil {
		return nil, err
	}
	profile := &dto.Profile{}
	err = provider.GetProfile(profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

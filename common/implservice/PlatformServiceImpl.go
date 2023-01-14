package implservice

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// PlatformServiceImpl ...
type PlatformServiceImpl struct {
	markup.Component `id:"PlatformService"`

	PSRs []service.PlatformServiceRegistry `inject:".PlatformServiceRegistry"`
}

func (inst *PlatformServiceImpl) _Impl() service.PlatformService {
	return inst
}

// Accept ...
func (inst *PlatformServiceImpl) Accept(p *dto.Platform) bool {
	return (p != nil)
}

// GetInfo ...
func (inst *PlatformServiceImpl) GetInfo(p *dto.Platform) (*dto.Platform, error) {

	arch := runtime.GOARCH
	os := runtime.GOOS

	if p == nil {
		p = &dto.Platform{}
	}

	p.Arch = strings.ToLower(arch)
	p.OS = strings.ToLower(os)

	dele, err := inst.findDelegate(p)
	if err != nil {
		return nil, err
	}

	return dele.GetInfo(p)
}

func (inst *PlatformServiceImpl) findDelegate(p *dto.Platform) (service.PlatformService, error) {

	if p == nil {
		return nil, fmt.Errorf("param is nil")
	}
	all := inst.PSRs
	for _, reg1 := range all {
		reg2 := reg1.GetRegistration()
		ser := reg2.Service
		if ser.Accept(p) {
			return ser, nil
		}
	}

	arch := p.Arch
	os := p.OS
	return nil, fmt.Errorf("unsupported platform, os:%v arch:%v", os, arch)
}

////////////////////////////////////////////////////////////////////////////////

func getRequiredEnvironmentValues(keys ...string) (map[string]string, error) {
	dst := make(map[string]string)
	for _, key := range keys {
		value := os.Getenv(key)
		if value == "" {
			return nil, fmt.Errorf("no required env:%v", key)
		}
		dst[key] = value
	}
	return dst, nil
}

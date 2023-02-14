// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package wpmserver

import (
	markup0x23084a "github.com/bitwormhole/starter/markup"
	implservice0x5a8f41 "github.com/bitwormhole/wpm/common/implservice"
	service0x3e063d "github.com/bitwormhole/wpm/server/service"
)

type pComPlatformServiceImpl struct {
	instance *implservice0x5a8f41.PlatformServiceImpl
	 markup0x23084a.Component `id:"PlatformService"`
	Providers []service0x3e063d.PlatformProviderRegistry `inject:".PlatformProviderRegistry"`
}


type pComProfileServiceImpl struct {
	instance *implservice0x5a8f41.ProfileServiceImpl
	 markup0x23084a.Component `id:"ProfileService"`
	PlatformService service0x3e063d.PlatformService `inject:"#PlatformService"`
}


type pComLinuxPlatformServiceImpl struct {
	instance *implservice0x5a8f41.LinuxPlatformServiceImpl
	 markup0x23084a.Component `class:"PlatformProviderRegistry"`
}


type pComWindowsPlatformServiceImpl struct {
	instance *implservice0x5a8f41.WindowsPlatformServiceImpl
	 markup0x23084a.Component `class:"PlatformProviderRegistry"`
}


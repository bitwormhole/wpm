// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package wpmclient

import (
	markup0x23084a "github.com/bitwormhole/starter/markup"
	client0x4f5e0e "github.com/bitwormhole/wpm/client"
	service0x3e063d "github.com/bitwormhole/wpm/server/service"
)

type pComClient struct {
	instance *client0x4f5e0e.Client
	 markup0x23084a.Component `class:"life"`
	IntentHandler client0x4f5e0e.IntentHandler `inject:"#IntentHandler"`
	PlatformService service0x3e063d.PlatformService `inject:"#PlatformService"`
	Protocol string `inject:"${wpm.server.protocol}"`
	Host string `inject:"${wpm.server.host}"`
	Port int `inject:"${wpm.server.port}"`
}


type pComIntentHandlerImpl struct {
	instance *client0x4f5e0e.IntentHandlerImpl
	 markup0x23084a.Component `id:"IntentHandler"`
}


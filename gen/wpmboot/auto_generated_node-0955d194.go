// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package wpmboot

import (
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
	boot0x8dc527 "github.com/bitwormhole/wpm/boot"
)

type pComInfoLoader struct {
	instance *boot0x8dc527.InfoLoader
	 markup0x23084a.Component `class:"life"`
	AC application0x67f6c5.Context `inject:"context"`
	DoBackupExe bool `inject:"${wpm.options.backup-this-exe}"`
	DoDebug bool `inject:"${wpm.options.debug}"`
	DoDump bool `inject:"${wpm.options.dump}"`
	DoLogOptions bool `inject:"${wpm.options.log-options}"`
	DoLogProps bool `inject:"${wpm.options.log-properties}"`
	DoRunWithGUI bool `inject:"${wpm.options.run-with-gui}"`
	DoRunWithServer bool `inject:"${wpm.options.run-with-server}"`
}


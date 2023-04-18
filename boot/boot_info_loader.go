package boot

import (
	"sort"
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

// InfoLoader ... 预先加载配置信息
type InfoLoader struct {
	markup.Component `class:"life"`

	AC application.Context `inject:"context"`

	DoBackupExe     bool `inject:"${wpm.options.backup-this-exe}"` // 是否备份当前EXE
	DoDebug         bool `inject:"${wpm.options.debug}"`           // 是否启用调试模式
	DoDump          bool `inject:"${wpm.options.dump}"`            // 是否备份转存数据
	DoLogOptions    bool `inject:"${wpm.options.log-options}"`     // 是否打印显示启动选项
	DoLogProps      bool `inject:"${wpm.options.log-properties}"`  // 是否打印显示属性列表
	DoRunWithGUI    bool `inject:"${wpm.options.run-with-gui}"`    // 是否显示图形界面
	DoRunWithServer bool `inject:"${wpm.options.run-with-server}"` // 是否启动后台服务
}

func (inst *InfoLoader) _Impl() application.LifeRegistry {
	return inst
}

// GetLifeRegistration ...
func (inst *InfoLoader) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnInit:   inst.init,
		Priority: 10000,
	}
}

func (inst *InfoLoader) init() error {
	if inst.DoLogOptions {
		inst.logOptions()
	}
	if inst.DoLogProps {
		inst.logProperties()
	}
	return nil
}

func (inst *InfoLoader) logProperties() {
	// const prefix = "wpm.options."
	all := inst.AC.GetProperties().Export(nil)
	keys := make([]string, 0)
	for key := range all {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	vlog.Info("properties:")
	for _, key := range keys {
		val := all[key]
		keyf := inst.formatKey(key)
		vlog.Info("    ", keyf, " = ", val)
	}
}

func (inst *InfoLoader) logOptions() {
	const prefix = "wpm.options."
	all := inst.AC.GetProperties().Export(nil)
	keys := make([]string, 0)
	for key := range all {
		if strings.HasPrefix(key, prefix) {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	vlog.Info("options:")
	for _, key := range keys {
		val := all[key]
		keyf := inst.formatKey(key)
		vlog.Info("    ", keyf, " = ", val)
	}
}

func (inst *InfoLoader) formatKey(key string) string {
	wantLen := 30
	haveLen := len(key)
	if haveLen < wantLen {
		builder := strings.Builder{}
		builder.WriteString(key)
		for ; haveLen < wantLen; haveLen++ {
			builder.WriteRune(' ')
		}
		return builder.String()
	}
	return key
}

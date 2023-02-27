package utils

import (
	"fmt"
	"strings"

	"github.com/bitwormhole/starter/application"
)

// CheckModuleVersion 检查各个模块的版本是否一致，如果不是就返回错误
func CheckModuleVersion(mods ...application.Module) error {

	want := ""
	msg := ""

	sb := strings.Builder{}
	sb.WriteString("modules:\n")

	for _, m := range mods {
		name := m.GetName()
		ver := m.GetVersion()
		rev := m.GetRevision()
		sb.WriteString(fmt.Sprintln(name, "@", ver, "(r", rev, ")"))
		have := ver
		if want == "" {
			want = have
		} else if want != have {
			msg = fmt.Sprintln("bad version, want:", want, " have:", have)
		}
	}

	if msg == "" {
		return nil
	}
	return fmt.Errorf("%v%v", msg, sb.String())
}

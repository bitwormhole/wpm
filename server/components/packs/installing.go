package packs

import (
	"context"
	"strings"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// InstallingContext 表示一个正在执行的安装任务
type InstallingContext struct {
	Context context.Context

	Action       string // [install|uninstall|upgrade|re-install]
	PackType     string
	Installation dxo.InstallationID

	Root     afs.Path // 系统根文件夹，不同的系统上具体位置可能不同
	PackDir  afs.Path // 解压后的包文件夹路径
	PackFile afs.Path // 压缩包文件的路径

	Pack *dto.SoftwarePackage
}

func (inst *InstallingContext) String() string {
	if inst == nil {
		return "nil"
	}
	p := inst.Pack
	b := strings.Builder{}
	b.WriteString("[packs.InstallingContext")
	b.WriteString(" action:" + inst.Action)
	if p != nil {
		b.WriteString(" content-type:" + p.ContentType)
		b.WriteString(" os:" + p.OS)
		b.WriteString(" arch:" + p.Arch)
	}
	b.WriteString("]")
	return b.String()
}

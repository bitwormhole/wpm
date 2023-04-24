package packs

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
)

// Manager 表示软件包管理器
// [inject:"#packs.Manager"]
type Manager interface {
	GetInstaller(ic *InstallingContext) (Installer, error)

	Install(c context.Context, pid dxo.SoftwarePackageID) error

	Uninstall(c context.Context, pid dxo.SoftwarePackageID) error
}

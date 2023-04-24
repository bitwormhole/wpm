package packs

// Installer 表示一个软件包安装器
type Installer interface {
	Accept(ic *InstallingContext) bool

	Install(ic *InstallingContext) error

	Uninstall(ic *InstallingContext) error
}

// InstallerRegistration 表示一个软件包安装器注册信息
type InstallerRegistration struct {
	Installer Installer
}

// InstallerRegistry 表示一个软件包安装器注册器
// [inject:".packs.InstallerRegistry"]
type InstallerRegistry interface {
	GetInstallerRegistration() *InstallerRegistration
}

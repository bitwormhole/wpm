package dxo

// SoftwarePackageState 表示软件包安装状态
type SoftwarePackageState string

// 定义软件包安装状态
const (
	SoftPackStateNA         SoftwarePackageState = "N/A"        // 无效
	SoftPackStateInstalled  SoftwarePackageState = "INSTALLED"  // 已安装
	SoftPackStateAvailable  SoftwarePackageState = "AVAILABLE"  // 可安装
	SoftPackStateNewVersion SoftwarePackageState = "NEWVERSION" // 可升级
)

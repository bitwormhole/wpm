package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// About ...
type About struct {
	Base

	Name      string `json:"name"`
	Title     string `json:"title"`
	Copyright string `json:"copyright"`
	Profile   string `json:"profile"`
	User      string `json:"user"`
	Home      string `json:"home"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
	WebURL    string `json:"weburl"`

	// the main module
	Module     string     `json:"module"`
	Version    string     `json:"version"`
	Revision   int        `json:"revision"`
	MainModule dto.Module `json:"main_module"`

	Modules []*dto.Module `json:"modules"`

	CheckUpdate *AboutCheckUpdate `json:"checkupdate"`
}

// AboutCheckUpdate  用于检查更新
type AboutCheckUpdate struct {

	// params
	Auto  bool `json:"auto"`  // 自动检查
	Force bool `json:"force"` // 必须检查

	// results
	Current       *dto.SoftwarePackage `json:"current"`         // 当前使用的版本
	Latest        *dto.SoftwarePackage `json:"latest"`          // 最新的版本
	Ignore        *dto.SoftwarePackage `json:"ignore"`          // 已忽略的版本
	HasNewVersion bool                 `json:"has_new_version"` // 检查结果
	HasIgnored    bool                 `json:"has_ignored"`     // 检查结果
}

package dto

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/starter-go/base/lang"
)

// SoftwarePackage 表示一个命令模板
type SoftwarePackage struct {
	ID  dxo.SoftwarePackageID  `json:"id"`
	URN dxo.SoftwarePackageURN `json:"urn"`

	Base

	Name       string `json:"name"`        // 简单文件名
	ModuleName string `json:"module_name"` // 包名称（=主模块.名称）
	FileName   string `json:"file_name"`   // 完整的文件名
	Namespace  string `json:"namespace"`   // 完整的文件名

	Icon        string `json:"icon"` // URL of icon
	Title       string `json:"title"`
	Description string `json:"description"`

	Revision int    `json:"revision"` // 包版次（=主模块.版次）
	Version  string `json:"version"`  // 包版本（=主模块.版本）
	OS       string `json:"os"`       // 操作系统
	Arch     string `json:"arch"`     // 处理器架构

	SHA256SUM   lang.Hex  `json:"sha256sum"`    // 包文件 sha-256
	Size        int64     `json:"size"`         // 包文件大小
	ContentType string    `json:"content_type"` // 包格式
	WebPageURL  string    `json:"web_page_url"` // 下载页面 URL
	DownloadURL string    `json:"download_url"` // 下载地址 URL
	ReleaseAt   lang.Time `json:"release_at"`   // 发布时间
	// ResourceURL string    `json:"resource_url"` // 包的资源下载 URL

	Installed bool                     `json:"installed"`
	State     dxo.SoftwarePackageState `json:"state"`
}

// SoftwareSet 表示一个软件集合，可能包含多个不同版本的包
type SoftwareSet struct {
	ID dxo.SoftwareSetID `json:"id"`

	// URN dxo.SoftwarePackageURN   `json:"urn"`
	// Base

	SoftwarePackage

	Packages []*SoftwarePackage `json:"packages"`
}

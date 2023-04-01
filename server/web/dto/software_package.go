package dto

import (
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/dxo"
)

// SoftwarePackage 表示一个命令模板
type SoftwarePackage struct {
	ID dxo.SoftwarePackageID `json:"id"`
	Base

	ModuleName string `json:"name"`        // 包名称（=主模块.名称）
	FileName   string `json:"full_name"`   // 完整的文件名
	AppName    string `json:"simple_name"` // 简单文件名

	Revision int    `json:"revision"` // 包版次（=主模块.版次）
	Version  string `json:"version"`  // 包版本（=主模块.版本）
	OS       string `json:"os"`       // 操作系统
	Arch     string `json:"arch"`     // 处理器架构

	SHA256SUM   util.Hex  `json:"sha256sum"`    // 包文件 sha-256
	Size        int64     `json:"size"`         // 包文件大小
	WebPageURL  string    `json:"web_page_url"` // 下载页面 URL
	DownloadURL string    `json:"download_url"` // 下载地址 URL
	MetaURL     string    `json:"meta_url"`     // 包的元数据下载 URL
	ReleaseAt   util.Time `json:"release_at"`   // 发布时间
}

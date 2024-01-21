package home

import (
	"context"

	"github.com/starter-go/afs"
)

// Service 提供一些关键的文件夹位置
type Service interface {

	// 获取由环境变量 WPM_HOME 指定的文件夹
	GetWPMHomeFolder() afs.Path

	// 获取 WPM 的数据文件夹
	GetWPMDataFolder() afs.Path

	// 获取当前用户的主目录
	GetUserHomeFolder(ctx context.Context) (afs.Path, error)

	// 获取当前用户的数据目录
	GetUserDataFolder(ctx context.Context) (afs.Path, error)
}

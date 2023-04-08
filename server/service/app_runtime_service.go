package service

import (
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/backup"
)

// AppRuntimeService ... 主要用来记录运行日志
type AppRuntimeService interface {

	// 查询日志
	// Find() error

	// 取当前 exe 文件的 sha-256
	GetAppHash() (util.Hex, error)

	// 从 runtime.json 中读取日志信息
	ReadStartupLogs() (*backup.StartupVO, error)
}

package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/web/vo"
)

// CheckUpdateService  检查更新 WPM 版本
type CheckUpdateService interface {

	// 检查更新的版本
	Check(ctx context.Context, o *vo.AboutCheckUpdate) error

	// 忽略指定的版本
	Ignore(ctx context.Context, o *vo.AboutCheckUpdate) error
}

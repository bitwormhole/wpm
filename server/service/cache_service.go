package service

import "github.com/bitwormhole/wpm/server/components/caches"

// CacheService 提供统一的缓存服务
type CacheService interface {
	GetManager() caches.Manager
}

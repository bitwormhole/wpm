package caches

import (
	"time"

	"github.com/starter-go/base/lang"
)

// Scope 表示缓存的作用域
type Scope int

// Scope 定义缓存的作用域
const (
	ScopeNone = iota
	ScopeItem
	ScopeClass
	ScopePool
	ScopeGlobal
	ScopeAll
)

// Service 表示缓存服务
type Service interface {
	GetPool() Pool
}

// Pool 表示缓存数据池
type Pool interface {
	// NewClass(name string) Class

	GetClass(name string) Class

	Clean(want *Want)
}

// OnLoadFunc 表示缓存数据加载器
type OnLoadFunc func(want *Want) (any, error)

// Class 表示一种缓存数据类型
type Class interface {
	Name() string

	UUID() lang.UUID

	// Loader() Loader

	MaxAge() time.Duration

	GetItem(id string) Item
}

// Item 表示一个缓存项
type Item interface {
	ID() string

	Class() Class

	CreatedAt() lang.Time

	UpdatedAt() lang.Time

	// 更新缓存数据
	Update()

	GetData(want *Want) (any, error)
}

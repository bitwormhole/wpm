package caches

// Manager 用来管理缓存
type Manager interface {

	// 清空缓存
	Clean()

	// 取缓存对象
	Get(p Provider) (any, error)
}

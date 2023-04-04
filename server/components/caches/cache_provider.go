package caches

// Provider 提供某个类的缓存
type Provider interface {

	// 新建缓存对象
	New() (any, error)

	// 缓存类型名称
	ClassName() string
}

// ProviderRegistration 表示缓存类的注册信息
type ProviderRegistration struct {
	ClassName string
	Enabled   bool
	Provider  Provider
}

// ProviderRegistry 缓存类的注册器
// [inject:".wpm-cache-provider"]
type ProviderRegistry interface {
	ListProviderRegistrations() []*ProviderRegistration
}

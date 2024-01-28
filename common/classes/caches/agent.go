package caches

import "context"

// Agent ...
type Agent struct {
	class Class
}

// NewLoader ...
func (inst *Agent) NewLoader(ctx context.Context) Loader {
	loader := &myLoader{}
	// loader.agent = inst
	loader.want.Context = ctx
	return loader
}

package intents

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/components/intents"
)

// FilterManagerImpl ...
type FilterManagerImpl struct {
	markup.Component `id:"wpm-intent-filter-manager"`

	FilterRegistryList []intents.FilterRegistry `inject:".wpm-intent-filter"`

	chain intents.FilterChain // cache of chain
}

func (inst *FilterManagerImpl) _Impl() intents.FilterManager {
	return inst
}

// Chain ...
func (inst *FilterManagerImpl) Chain() intents.FilterChain {
	ch := inst.chain
	if ch != nil {
		return ch
	}
	builder := &intents.FilterChainBuilder{}
	builder.AddRegistry(inst.FilterRegistryList...)
	ch = builder.Create()
	inst.chain = ch
	return ch
}

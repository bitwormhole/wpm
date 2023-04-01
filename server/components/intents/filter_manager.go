package intents

// FilterManager ... [inject:"#wpm-intent-filter-manager"]
type FilterManager interface {
	Chain() FilterChain
}

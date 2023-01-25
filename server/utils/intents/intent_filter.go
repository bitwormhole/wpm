package intents

import (
	"github.com/bitwormhole/wpm/server/web/dto"
)

// Filter ...
type Filter interface {
	Filter(i *dto.Intent) (*dto.Intent, error)
}

// FilterRegistry ... [inject:".intent-filter-registry"]
type FilterRegistry interface {
	GetRegistrationList() []*FilterRegistration
}

// FilterRegistration ...
type FilterRegistration struct {
	Filter Filter
	Order  int
}

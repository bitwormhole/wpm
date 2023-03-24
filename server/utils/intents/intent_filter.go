package intents

import (
	"context"

	"github.com/bitwormhole/wpm/server/web/dto"
)

// Filter ...
type Filter interface {
	Filter(c context.Context, i *dto.Intent) (*dto.Intent, error)
}

// FilterRegistry ... [inject:".intent-filter-registry"]
type FilterRegistry interface {
	GetRegistrationList() []*FilterRegistration
}

// FilterRegistration ...
type FilterRegistration struct {
	Filter  Filter
	Order   int
	Enabled bool
}

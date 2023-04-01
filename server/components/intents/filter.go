package intents

import (
	"context"

	"github.com/bitwormhole/wpm/server/web/dto"
)

// Filter ...
type Filter interface {
	Handle(c context.Context, i *dto.Intent, next FilterChain) error
}

// FilterChain ...
type FilterChain interface {
	Handle(c context.Context, i *dto.Intent) error
}

// FilterRegistration ...
type FilterRegistration struct {
	Filter  Filter
	Order   int
	Enabled bool
}

// FilterRegistry ... [inject:".wpm-intent-filter"]
type FilterRegistry interface {
	GetFilterRegistrationList() []*FilterRegistration
}

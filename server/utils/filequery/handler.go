package filequery

import "github.com/bitwormhole/wpm/server/web/vo"

type Handler interface {
	Accept(q *vo.FileQuery) bool
	Query(q *vo.FileQuery) (*vo.FileQuery, error)
}

// HandlerRegistry [inject:".filequery-handler-registry"]
type HandlerRegistry interface {
	ListRegistrations() []*HandlerRegistration
}

type HandlerRegistration struct {
	Enabled bool
	Handler Handler
}

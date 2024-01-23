package ifilequery

import (
	"context"

	"github.com/bitwormhole/wpm/server/classes/filequery"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/starter-go/afs"
)

// ServiceImpl ...
type ServiceImpl struct {

	//starter:component
	////// starter:as("#")
	_as func(filequery.Service) //starter:as("#")

	FS afs.FS //starter:inject("#")

}

func (inst *ServiceImpl) _impl() filequery.Service {
	return inst
}

// Query ...
func (inst *ServiceImpl) Query(ctx context.Context, want *vo.FileQuery) (*vo.FileQuery, error) {
	finder := &fsDirItemsFinder{
		parent: inst,
	}
	return finder.find(want)
}

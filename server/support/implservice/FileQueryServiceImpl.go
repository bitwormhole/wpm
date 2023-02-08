package implservice

import (
	"context"
	"errors"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// FileQueryServiceImpl ...
type FileQueryServiceImpl struct {
	markup.Component `id:"FileQueryService"`

	// FileQueryDAO dao.FileQueryDAO `inject:"#FileQueryDAO"`
}

func (inst *FileQueryServiceImpl) _Impl() service.FileQueryService {
	return inst
}

// Query ...
func (inst *FileQueryServiceImpl) Query(ctx context.Context, q *vo.FileQuery) (*vo.FileQuery, error) {

	// src, err := inst.FileQueryDAO.ListAll()
	// if err != nil {
	// 	return nil, err
	// }
	// dst := make([]*dto.FileQuery, 0)
	// for _, item1 := range src {
	// 	item2, err := inst.entity2dto(item1)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	dst = append(dst, item2)
	// }
	// return dst, nil

	return nil, errors.New("no impl")
}

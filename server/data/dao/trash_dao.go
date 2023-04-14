package dao

import "github.com/bitwormhole/wpm/server/data/entity"

// TrashDAO 用来存取已经被软删除的对象
type TrashDAO interface {
	ListAll() ([]*entity.Holder, error)

	Delete(h ...*entity.Holder) error

	Recover(h ...*entity.Holder) error
}

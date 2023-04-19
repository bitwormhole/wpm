package service

import "github.com/bitwormhole/wpm/server/web/vo"

// OptionService ...
type OptionService interface {
	GetOptions(o *vo.Option) error
}

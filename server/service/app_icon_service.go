package service

import "github.com/bitwormhole/wpm/server/web/dto"

// AppIconService ...
type AppIconService interface {
	FillWithIconURL(o1 *dto.Executable) error
}

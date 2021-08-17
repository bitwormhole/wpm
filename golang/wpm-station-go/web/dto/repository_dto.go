package dto

import "github.com/bitwormhole/wpm/web/base"

// Repository 仓库DTO
type Repository struct {
	base.DTO
	DisplayName string
	Path        string
	Ready       bool
}

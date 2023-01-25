package service

import "github.com/bitwormhole/wpm/server/web/dto"

// ProfileService ...
type ProfileService interface {
	GetProfile() (*dto.Profile, error)
}

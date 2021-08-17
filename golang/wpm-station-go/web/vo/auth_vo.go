package vo

import "github.com/bitwormhole/wpm/web/base"

// Auth 身份认证VO
type Auth struct {
	base.VO

	UserName   string
	UserSecret string
	Mechanism  string

	Authorized bool
}

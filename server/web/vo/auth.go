package vo

// Auth 身份认证VO
type Auth struct {
	Base

	UserName   string
	UserSecret string
	Mechanism  string

	Authorized bool
}

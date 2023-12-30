package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// MIMEType ...
type MIMEType struct {
	Base

	Types []*dto.MIMEType `json:"types"`
}

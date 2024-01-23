package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// MIMEType ...
type MIMEType struct {
	Base

	Types []*dto.MIMEType `json:"types"`
}

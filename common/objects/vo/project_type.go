package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// ContentType ...
type ContentType struct {
	Base

	ContentTypes []*dto.ContentType `json:"content_types"`
}

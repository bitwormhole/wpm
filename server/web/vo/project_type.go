package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// ContentType ...
type ContentType struct {
	Base

	ContentTypes []*dto.ContentType `json:"content_types"`
}

package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// ContentTypeImport ...
type ContentTypeImport struct {
	Base

	ContentTypes []*dto.ContentType `json:"content_types"`
}

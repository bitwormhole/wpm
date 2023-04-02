package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// ContentTypeImport ...
type ContentTypeImport struct {
	Base

	ContentTypes []*dto.ContentType `json:"content_types"`
}

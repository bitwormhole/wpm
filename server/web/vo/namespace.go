package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// Namespace ...
type Namespace struct {
	Base

	Namespaces []*dto.Namespace `json:"namespaces"`
}

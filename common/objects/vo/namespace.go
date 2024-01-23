package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// Namespace ...
type Namespace struct {
	Base

	Namespaces []*dto.Namespace `json:"namespaces"`
}

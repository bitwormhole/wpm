package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// Setup ...
type Setup struct {
	Base

	SetupRequired bool `json:"setup_required"`

	SetupItems []*dto.Setup `json:"setup_items"`
}

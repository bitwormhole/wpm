package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// Statistic ...
type Statistic struct {
	Base

	Statistics dto.Statistic `json:"statistics"`
}

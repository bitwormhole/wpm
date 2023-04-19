package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// Statistic ...
type Statistic struct {
	Base

	Statistics dto.Statistic `json:"statistics"`
}

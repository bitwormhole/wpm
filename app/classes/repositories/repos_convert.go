package repositories

import (
	"fmt"

	"github.com/bitwormhole/wpm/app/web/dto"
)

// Convertor ...
type Convertor struct{}

// ConvertD2L ...
func (inst *Convertor) ConvertD2L(src *dto.LocalRepository) (*dto.Location, error) {

	return nil, fmt.Errorf("no impl ")
}

// ConvertL2D ...
func (inst *Convertor) ConvertL2D(src *dto.Location) (*dto.LocalRepository, error) {

	return nil, fmt.Errorf("no impl ")
}

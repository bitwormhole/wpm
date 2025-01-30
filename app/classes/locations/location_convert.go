package locations

import (
	"github.com/bitwormhole/wpm/app/data/entity"
	"github.com/bitwormhole/wpm/app/web/dto"
	"github.com/starter-go/security-gorm/rbacdb"
)

// ConvertD2E ...
func ConvertD2E(src *dto.Location) *entity.Location {

	dst := new(entity.Location)
	dst.ID = src.ID
	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)

	dst.Label = src.Label
	dst.Description = src.Description
	dst.Tags = src.Tags

	return dst
}

// ConvertE2D ...
func ConvertE2D(src *entity.Location) *dto.Location {

	dst := new(dto.Location)
	dst.ID = src.ID
	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)

	dst.Label = src.Label
	dst.Description = src.Description
	dst.Tags = src.Tags

	return dst
}

// ConvertListE2D ...
func ConvertListE2D(src []*entity.Location) []*dto.Location {
	dst := make([]*dto.Location, 0)
	for _, item1 := range src {
		item2 := ConvertE2D(item1)
		dst = append(dst, item2)
	}
	return dst
}

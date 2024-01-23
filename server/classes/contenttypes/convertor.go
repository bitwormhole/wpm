package contenttypes

import (
	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/starter-go/security-gorm/rbacdb"
)

// Convertor ...
type Convertor struct{}

// NewConvertor ...
func NewConvertor() *Convertor {
	return new(Convertor)
}

// ConvertListE2D ...
func (inst *Convertor) ConvertListE2D(src []*entity.ContentType) []*dto.ContentType {
	dst := make([]*dto.ContentType, 0)
	for _, item1 := range src {
		item2 := inst.ConvertE2D(item1)
		dst = append(dst, item2)
	}
	return dst
}

// ConvertE2D ...
func (inst *Convertor) ConvertE2D(src *entity.ContentType) *dto.ContentType {
	dst := new(dto.ContentType)
	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)
	dst.ID = src.ID

	//todo ....

	return dst
}

// ConvertD2E ...
func (inst *Convertor) ConvertD2E(src *dto.ContentType) *entity.ContentType {
	dst := new(entity.ContentType)
	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)
	dst.ID = src.ID

	//todo ....

	return dst
}

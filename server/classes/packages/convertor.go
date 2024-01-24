package packages

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
func (inst *Convertor) ConvertListE2D(src []*entity.SoftwarePackage) []*dto.SoftwarePackage {
	dst := make([]*dto.SoftwarePackage, 0)
	for _, item1 := range src {
		item2 := inst.ConvertE2D(item1)
		dst = append(dst, item2)
	}
	return dst
}

// ConvertE2D ...
func (inst *Convertor) ConvertE2D(src *entity.SoftwarePackage) *dto.SoftwarePackage {
	dst := new(dto.SoftwarePackage)
	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)
	dst.ID = src.ID

	//todo ....

	return dst
}

// ConvertD2E ...
func (inst *Convertor) ConvertD2E(src *dto.SoftwarePackage) *entity.SoftwarePackage {
	dst := new(entity.SoftwarePackage)
	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)
	dst.ID = src.ID

	//todo ....

	return dst
}
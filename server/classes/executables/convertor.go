package executables

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
func (inst *Convertor) ConvertListE2D(src []*entity.Executable) []*dto.Executable {
	dst := make([]*dto.Executable, 0)
	for _, item1 := range src {
		item2 := inst.ConvertE2D(item1)
		dst = append(dst, item2)
	}
	return dst
}

// ConvertE2D ...
func (inst *Convertor) ConvertE2D(src *entity.Executable) *dto.Executable {
	dst := new(dto.Executable)
	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)
	dst.ID = src.ID

	//todo ....
	dst.URN = src.URN
	dst.Name = src.Name
	dst.Aliases = src.Aliases
	dst.Namespace = src.Namespace
	dst.Title = src.Title
	dst.Description = src.Description
	dst.IconURL = src.IconURL

	dst.Path = src.Path
	dst.Size = src.Size
	dst.SHA256SUM = src.SHA256SUM
	dst.OpenWithPriority = src.OpenWithPriority

	dst.OS = src.OS
	dst.Arch = src.Arch
	dst.Version = src.Version

	return dst
}

// ConvertD2E ...
func (inst *Convertor) ConvertD2E(src *dto.Executable) *entity.Executable {
	dst := new(entity.Executable)
	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)
	dst.ID = src.ID

	//todo ....
	dst.URN = src.URN
	dst.Name = src.Name
	dst.Aliases = src.Aliases
	dst.Namespace = src.Namespace
	dst.Title = src.Title
	dst.Description = src.Description
	dst.IconURL = src.IconURL

	dst.Path = src.Path
	dst.Size = src.Size
	dst.SHA256SUM = src.SHA256SUM
	dst.OpenWithPriority = src.OpenWithPriority

	dst.OS = src.OS
	dst.Arch = src.Arch
	dst.Version = src.Version

	return dst
}

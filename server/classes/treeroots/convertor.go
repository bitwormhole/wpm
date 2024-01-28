package treeroots

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
func (inst *Convertor) ConvertListE2D(src []*entity.TreeRoot) []*dto.TreeRoot {
	dst := make([]*dto.TreeRoot, 0)
	for _, item1 := range src {
		item2 := inst.ConvertE2D(item1)
		dst = append(dst, item2)
	}
	return dst
}

// ConvertE2D ...
func (inst *Convertor) ConvertE2D(src *entity.TreeRoot) *dto.TreeRoot {
	dst := new(dto.TreeRoot)
	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)
	dst.ID = src.ID

	//todo ....
	dst.Path = src.Path
	dst.DotGitPath = src.DotGitPath
	dst.Class = src.Class

	return dst
}

// ConvertD2E ...
func (inst *Convertor) ConvertD2E(src *dto.TreeRoot) *entity.TreeRoot {
	dst := new(entity.TreeRoot)
	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)
	dst.ID = src.ID

	//todo ....
	dst.Path = src.Path
	dst.DotGitPath = src.DotGitPath
	dst.Class = src.Class

	return dst
}

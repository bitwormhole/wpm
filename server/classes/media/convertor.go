package media

import (
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/starter-go/security-gorm/rbacdb"
)

// Convertor ...
type Convertor struct{}

// NewConvertor ...
func NewConvertor() *Convertor {
	return new(Convertor)
}

// ConvertListE2D ...
func (inst *Convertor) ConvertListE2D(src []*entity.Media) []*dto.Media {
	dst := make([]*dto.Media, 0)
	for _, item1 := range src {
		item2 := inst.ConvertE2D(item1)
		dst = append(dst, item2)
	}
	return dst
}

// ConvertE2D ...
func (inst *Convertor) ConvertE2D(src *entity.Media) *dto.Media {
	dst := new(dto.Media)
	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)
	dst.ID = src.ID

	//todo ....
	dst.Name = src.Name
	dst.FileSize = src.ContentLength
	dst.ContentType = src.ContentType
	dst.SHA256SUM = src.SHA256SUM
	dst.Label = src.Label
	dst.Source = src.Source
	dst.URL = src.URL
	dst.Bucket = src.Bucket

	return dst
}

// ConvertD2E ...
func (inst *Convertor) ConvertD2E(src *dto.Media) *entity.Media {
	dst := new(entity.Media)
	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)
	dst.ID = src.ID

	//todo ....
	dst.Name = src.Name
	dst.ContentLength = src.FileSize
	dst.ContentType = src.ContentType
	dst.SHA256SUM = src.SHA256SUM
	dst.Label = src.Label
	dst.Source = src.Source
	dst.URL = src.URL
	dst.Bucket = src.Bucket

	return dst
}

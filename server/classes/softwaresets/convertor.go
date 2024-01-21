package softwaresets

// Convertor ...
type Convertor struct{}

// NewConvertor ...
func NewConvertor() *Convertor {
	return new(Convertor)
}

// // ConvertListE2D ...
// func (inst *Convertor) ConvertListE2D(src []*entity.SoftwareSet) []*dto.SoftwareSet {
// 	dst := make([]*dto.SoftwareSet, 0)
// 	for _, item1 := range src {
// 		item2 := inst.ConvertE2D(item1)
// 		dst = append(dst, item2)
// 	}
// 	return dst
// }

// // ConvertE2D ...
// func (inst *Convertor) ConvertE2D(src *entity.SoftwareSet) *dto.SoftwareSet {
// 	dst := new(dto.SoftwareSet)
// 	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)
// 	dst.ID = src.ID

// 	//todo ....

// 	return dst
// }

// // ConvertD2E ...
// func (inst *Convertor) ConvertD2E(src *dto.SoftwareSet) *entity.SoftwareSet {
// 	dst := new(entity.SoftwareSet)
// 	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)
// 	dst.ID = src.ID

// 	//todo ....

// 	return dst
// }

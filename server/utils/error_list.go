package utils

// ErrorList 表示一个可以包含多个错误的列表
type ErrorList struct {
	list []error
}

// Append 添加一个错误到列表末尾
func (inst *ErrorList) Append(err error) {
	if err == nil {
		return
	}
	inst.list = append(inst.list, err)
}

// First  取最第一个错误
func (inst *ErrorList) First() error {
	size := len(inst.list)
	if size > 0 {
		return inst.list[0]
	}
	return nil
}

// Last  取最后一个错误
func (inst *ErrorList) Last() error {
	size := len(inst.list)
	if size > 0 {
		return inst.list[size-1]
	}
	return nil
}

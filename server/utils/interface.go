package utils

import "reflect"

// InterfaceIsNil 判断 i 是否为空
func InterfaceIsNil(i any) bool {
	defer func() {
		recover()
	}()
	vi := reflect.ValueOf(i)
	return vi.IsNil()
}

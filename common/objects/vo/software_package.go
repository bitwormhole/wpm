package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// SoftwarePackage ... 软件包信息
type SoftwarePackage struct {
	Base

	Packages []*dto.SoftwarePackage `json:"packages"`
}

// SoftwareSet ... 软件集合信息
type SoftwareSet struct {
	Base

	Sets []*dto.SoftwareSet `json:"sets"`
}

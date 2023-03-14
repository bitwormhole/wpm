package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// SoftwarePackage ... 软件包信息
type SoftwarePackage struct {
	Base

	Packages []*dto.SoftwarePackage `json:"packages"`
}

package vo

import "github.com/bitwormhole/wpm/common/objects/dto"

// // OnlineHead ...
// type OnlineHead struct {
// }

// Online ...
type Online struct {
	Base

	Sources         []*dto.Namespace       `json:"sources"`
	Packages        []*dto.SoftwarePackage `json:"packages"`
	Mediae          []*dto.Media           `json:"mediae"`
	Executables     []*dto.Executable      `json:"executables"`
	ContentTypes    []*dto.ContentType     `json:"contenttypes"`
	IntentTemplates []*dto.IntentTemplate  `json:"intenttemplates"`
}

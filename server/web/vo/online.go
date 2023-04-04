package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// Online ...
type Online struct {
	Base

	OnlineURL string `json:"online"`

	Namespaces      []*dto.Namespace       `json:"namespaces"`
	Packages        []*dto.SoftwarePackage `json:"packages"`
	Mediae          []*dto.Media           `json:"mediae"`
	Executables     []*dto.Executable      `json:"executables"`
	ContentTypes    []*dto.ContentType     `json:"contenttypes"`
	IntentTemplates []*dto.IntentTemplate  `json:"intenttemplates"`
}

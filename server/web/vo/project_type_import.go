package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// ProjectTypeImport ...
type ProjectTypeImport struct {
	Base

	ProjectTypes []*dto.ProjectType `json:"project_types"`
}

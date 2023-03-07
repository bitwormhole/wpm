package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// ProjectType ...
type ProjectType struct {
	Base

	ProjectTypes []*dto.ProjectType `json:"project_types"`
}

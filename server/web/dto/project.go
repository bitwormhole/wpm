package dto

import "github.com/bitwormhole/wpm/server/data/dxo"

// Project ...
type Project struct {
	ID dxo.ProjectID `json:"id"`
	Base

	Name            string                `json:"name"`
	Description     string                `json:"description"`
	ProjectType     string                `json:"project_type"`
	Path            string                `json:"path"`
	OwnerRepository dxo.LocalRepositoryID `json:"owner_repository"`
	Group           dxo.ProjectGroupName  `json:"group"`
	State           dxo.FileState         `json:"state"`
	Tags            dxo.StringList        `json:"tags"`
}

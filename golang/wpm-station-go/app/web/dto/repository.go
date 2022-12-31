package dto

import "github.com/bitwormhole/wpm/app/data/dxo"

// Repository 仓库DTO
type Repository struct {
	ID dxo.RepositoryID `json:"id"`
	Base

	Name        string                  `json:"name"`
	DisplayName string                  `json:"label"`
	Description string                  `json:"description"`
	Path        string                  `json:"path"`
	Ready       bool                    `json:"ready"`
	Group       dxo.RepositoryGroupName `json:"group"`
	State       dxo.FileState           `json:"state"`
	Tags        dxo.StringList          `json:"tags"`
}

// RepositoryGroup 是 Repository 的分组
type RepositoryGroup struct {
	ID dxo.RepositoryGroupID `json:"id"`
	Base

	Name string `json:"name"`
}

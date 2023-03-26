package dto

import "github.com/bitwormhole/wpm/server/data/dxo"

// RepositoryBase 仓库DTO
type RepositoryBase struct {
	Base

	Name        string                  `json:"name"`
	DisplayName string                  `json:"label"`
	Description string                  `json:"description"`
	IconURL     string                  `json:"icon"`
	Ready       bool                    `json:"ready"`
	Bare        bool                    `json:"bare"`
	Group       dxo.RepositoryGroupName `json:"group"`
	State       dxo.FileState           `json:"state"`
	Tags        dxo.StringList          `json:"tags"`
}

// LocalRepository ...
type LocalRepository struct {
	ID dxo.LocalRepositoryID `json:"id"`

	RepositoryBase

	DotGitPath     string `json:"dot_git_path"`
	RepositoryPath string `json:"repository_path"`
	WorkingPath    string `json:"workspace_path"`
	ConfigFile     string `json:"config_file_path"`

	Path     string            `json:"path"` // this.Path == this.RepositoryPath
	Location dxo.LocationID    `json:"location"`
	Class    dxo.LocationClass `json:"location_class"`

	Projects []*Project `json:"projects"`
}

// MainRepository ...
type MainRepository struct {
	RepositoryBase

	ID dxo.MainRepositoryID `json:"id"`

	Path string `json:"path"`
}

// RemoteRepository ...
type RemoteRepository struct {
	RepositoryBase

	ID dxo.RemoteRepositoryID `json:"id"`

	URL string `json:"url"`
}

// RepositoryGroup 是 Repository 的分组
type RepositoryGroup struct {
	ID dxo.RepositoryGroupID `json:"id"`
	Base

	Name string `json:"name"`
}

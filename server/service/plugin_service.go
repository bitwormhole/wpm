package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// SoftwarePackageService ...
type SoftwarePackageService interface {
	Find(ctx context.Context, id dxo.SoftwarePackageID) (*dto.SoftwarePackage, error)

	ListAll(ctx context.Context) ([]*dto.SoftwarePackage, error)

	ListByModuleName(ctx context.Context, moduleName string) ([]*dto.SoftwarePackage, error)

	Insert(ctx context.Context, o *dto.SoftwarePackage) (*dto.SoftwarePackage, error)

	UpdateItem(ctx context.Context, id dxo.SoftwarePackageID, o *dto.SoftwarePackage) (*dto.SoftwarePackage, error)

	UpdateList(ctx context.Context) error

	Remove(ctx context.Context, id dxo.SoftwarePackageID) error

	Install(ctx context.Context, id dxo.SoftwarePackageID) error

	Uninstall(ctx context.Context, id dxo.SoftwarePackageID) error
}

// SoftwareSetService ...
type SoftwareSetService interface {
	GetOne(ctx context.Context, id dxo.SoftwarePackageID) (*dto.SoftwareSet, error)

	ListAll(ctx context.Context) ([]*dto.SoftwareSet, error)

	Install(ctx context.Context, ss *dto.SoftwareSet) error

	ReInstall(ctx context.Context, ss *dto.SoftwareSet) error

	Uninstall(ctx context.Context, ss *dto.SoftwareSet) error

	Upgrade(ctx context.Context, ss *dto.SoftwareSet) error
}

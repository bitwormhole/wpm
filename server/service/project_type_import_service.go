package service

import "context"

// ProjectTypeImportService ...
type ProjectTypeImportService interface {
	ImportTypesFromPreset(ctx context.Context) error
}

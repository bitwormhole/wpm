package service

import (
	"context"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// IntentService ...
type IntentService interface {
	Run(ctx context.Context, o *dto.Intent) (*dto.Intent, error)
}

// IntentTemplateService ...
type IntentTemplateService interface {
	Find(ctx context.Context, id dxo.IntentTemplateID) (*dto.IntentTemplate, error)

	ListAll(ctx context.Context) ([]*dto.IntentTemplate, error)

	ListMacroProperties(ctx context.Context) (map[string]string, error)

	// find items by Action & Executable & Target
	ListByAET(ctx context.Context, sel *dto.IntentTemplate) ([]*dto.IntentTemplate, error)

	Insert(ctx context.Context, o *dto.IntentTemplate) (*dto.IntentTemplate, error)

	Update(ctx context.Context, id dxo.IntentTemplateID, o *dto.IntentTemplate) (*dto.IntentTemplate, error)

	Remove(ctx context.Context, id dxo.IntentTemplateID) error

	ImportPreset(ctx context.Context) error
}

// IntentHandlerService ...
type IntentHandlerService interface {
	HandleIntent(i *dto.Intent) error
}

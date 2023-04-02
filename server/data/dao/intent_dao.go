package dao

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
)

// IntentTemplateDAO ...
type IntentTemplateDAO interface {
	Find(id dxo.IntentTemplateID) (*entity.IntentTemplate, error)

	ListAll() ([]*entity.IntentTemplate, error)

	// find items by Action & Executable & Target & Type
	ListBySelector(selector *entity.IntentTemplate) ([]*entity.IntentTemplate, error)

	Insert(o *entity.IntentTemplate) (*entity.IntentTemplate, error)

	Update(id dxo.IntentTemplateID, o *entity.IntentTemplate) (*entity.IntentTemplate, error)

	Remove(id dxo.IntentTemplateID) error
}

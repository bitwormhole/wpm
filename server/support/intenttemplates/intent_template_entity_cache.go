package intenttemplates

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
)

// ImpIntentTemplateEntityCache ...
type ImpIntentTemplateEntityCache struct {
	markup.Component `id:"IntentTemplateEntityCache"`

	DAO dao.IntentTemplateDAO `inject:"#IntentTemplateDAO"`
}

func (inst *ImpIntentTemplateEntityCache) _Impl() service.IntentTemplateEntityCache {
	return inst
}

// ListTemplates ...
func (inst *ImpIntentTemplateEntityCache) ListTemplates() ([]*entity.IntentTemplate, error) {
	return inst.DAO.ListAll()
}

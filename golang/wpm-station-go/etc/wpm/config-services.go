package etcwpm

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/data/dao"
	"github.com/bitwormhole/wpm/service"
)

type theRepositoryService struct {
	markup.Component
	instance *service.RepositoryServiceImpl `id:"repository-service"`
	Store    *dao.Store                     `inject:"#store"`
}

package etcwpm

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/data/dao"
)

type theStore struct {
	markup.Component
	instance        *dao.Store       `id:"store"`
	RepositoryDAO   dao.Repository   `inject:"#data-repository-dao"`
	SystemConfigDAO dao.SystemConfig `inject:"#data-system-config-dao"`
	UserConfigDAO   dao.UserConfig   `inject:"#data-user-config-dao"`
}

type theRepositoryDAO struct {
	markup.Component
	instance *dao.RepositoryImpl `id:"data-repository-dao"`
}

type theSystemConfigDAO struct {
	markup.Component
	instance *dao.SystemConfigImpl `id:"data-system-config-dao"`
}

type theUserConfigDAO struct {
	markup.Component
	instance *dao.UserConfigImpl `id:"data-user-config-dao"`
}

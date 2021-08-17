package etcwpm

import (
	application "github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/service"
	"github.com/bitwormhole/wpm/web/controller"
)

type theRestControllerDebug struct {
	markup.Component
	instance   *controller.DebugController `class:"rest-controller"`
	AppContext application.Context         `inject:"context"`
}

type theRestControllerAuth struct {
	markup.Component
	instance *controller.AuthController `class:"rest-controller"`
}

type theRestControllerRepository struct {
	markup.Component
	instance    *controller.RepositoryController `class:"rest-controller"`
	RepoService service.RepositoryService        `inject:"#repository-service"`
}

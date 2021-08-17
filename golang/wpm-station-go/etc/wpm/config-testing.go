package etcwpm

import (
	"github.com/bitwormhole/gitlib/repository"
	application "github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/web/controller"
)

type theTestGitController struct {
	markup.Component
	instance   *controller.TestGitController `class:"rest-controller"`
	AppContext application.Context           `inject:"context"`
	RM         repository.Manager            `inject:"#git-repository-manager"`
}

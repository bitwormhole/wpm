package code

import (
	"github.com/bitwormhole/wpm/server/classes/contenttypes"
	"github.com/bitwormhole/wpm/server/classes/executables"
	"github.com/bitwormhole/wpm/server/classes/intenttemplates"
	"github.com/bitwormhole/wpm/server/classes/locations"
	"github.com/bitwormhole/wpm/server/classes/media"
	"github.com/bitwormhole/wpm/server/classes/packages"
	"github.com/bitwormhole/wpm/server/classes/projects"
	"github.com/bitwormhole/wpm/server/classes/repositories"
	"github.com/bitwormhole/wpm/server/classes/settings"
	"github.com/bitwormhole/wpm/server/classes/softwaresets"
	"github.com/starter-go/application"
)

// DebugAllDAOs ...
type DebugAllDAOs struct {

	//starter:component

	ContentTypeDAO     contenttypes.DAO                 //starter:inject("#")
	ExecutableDAO      executables.DAO                  //starter:inject("#")
	IntentTemplateDAO  intenttemplates.DAO              //starter:inject("#")
	LocationDAO        locations.DAO                    //starter:inject("#")
	MediaDAO           media.DAO                        //starter:inject("#")
	ProjectDAO         projects.DAO                     //starter:inject("#")
	LocalRepoDAO       repositories.LocalRepositoryDAO  //starter:inject("#")
	RemoteRepoDAO      repositories.RemoteRepositoryDAO //starter:inject("#")
	SettingDAO         settings.DAO                     //starter:inject("#")
	SoftwarePackageDAO packages.DAO                     //starter:inject("#")
	SoftwareSetDAO     softwaresets.DAO                 //starter:inject("#")

}

func (inst *DebugAllDAOs) _impl() application.Lifecycle {
	return inst
}

// Life ...
func (inst *DebugAllDAOs) Life() *application.Life {
	return &application.Life{
		OnLoop: inst.run,
	}
}

func (inst *DebugAllDAOs) run() error {

	inst.ContentTypeDAO.List(nil, &contenttypes.Query{})
	inst.ExecutableDAO.List(nil, &executables.Query{})

	return nil
}

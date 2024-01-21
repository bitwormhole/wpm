package wpm

import (
	"github.com/bitwormhole/gitlib/modules/gitlib"
	"github.com/bitwormhole/wpm"
	"github.com/bitwormhole/wpm/gen/agent4wpm"
	"github.com/bitwormhole/wpm/gen/boot4wpm"
	"github.com/bitwormhole/wpm/gen/cli4wpm"
	"github.com/bitwormhole/wpm/gen/common4wpm"
	"github.com/bitwormhole/wpm/gen/gui4wpm"
	"github.com/bitwormhole/wpm/gen/main4wpm"
	"github.com/bitwormhole/wpm/gen/server4wpm"
	"github.com/bitwormhole/wpm/gen/test4wpm"
	"github.com/starter-go/application"
	"github.com/starter-go/httpagent/modules/httpagent"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/module-gorm-mysql/modules/mysql"
	"github.com/starter-go/module-gorm-sqlserver/modules/sqlserver"
	"github.com/starter-go/security-gin-gorm/modules/securitygingorm"
	"github.com/starter-go/security/modules/security"
	"github.com/starter-go/starter"
)

// Module  the MAIN module
func Module() application.Module {
	mb := wpm.NewMainModule()
	mb.Components(main4wpm.ExportComponents)
	mb.Depend(ModuleForCommon())
	return mb.Create()
}

// AgentModule  ...
func AgentModule() application.Module {
	mb := wpm.NewAgentModule()
	mb.Components(agent4wpm.ExportComponents)
	mb.Depend(starter.Module())
	mb.Depend(httpagent.Module())
	mb.Depend(ModuleForCommon())
	return mb.Create()
}

// ServerModule  ...
func ServerModule() application.Module {
	mb := wpm.NewServerModule()
	mb.Components(server4wpm.ExportComponents)

	mb.Depend(libgin.Module())
	mb.Depend(libgorm.Module())
	mb.Depend(security.Module())
	mb.Depend(securitygingorm.Module())
	mb.Depend(gitlib.Module())
	mb.Depend(httpagent.Module())
	mb.Depend(ModuleForCommon())

	mb.Depend(mysql.Module())
	mb.Depend(sqlserver.Module())

	return mb.Create()
}

// ModuleForGUI ...
func ModuleForGUI() application.Module {
	mb := wpm.NewGuiModule()
	mb.Components(gui4wpm.ExportComponents)
	mb.Depend(ModuleForCommon())
	return mb.Create()
}

// ModuleForCLI ...
func ModuleForCLI() application.Module {
	mb := wpm.NewCliModule()
	mb.Components(cli4wpm.ExportComponents)
	mb.Depend(ModuleForCommon())
	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := wpm.NewTestModule()
	mb.Components(test4wpm.ExportComponents)
	mb.Depend(ServerModule())
	mb.Depend(ModuleForCommon())
	return mb.Create()
}

// ModuleForBoot  ...
func ModuleForBoot() application.Module {
	mb := wpm.NewBootModule()
	mb.Components(boot4wpm.ExportComponents)
	mb.Depend(ModuleForCommon())
	return mb.Create()
}

// ModuleForCommon ...
func ModuleForCommon() application.Module {
	mb := wpm.NewCommonModule()
	mb.Components(common4wpm.ExportComponents)
	return mb.Create()
}

package wpm

import (
	"github.com/bitwormhole/wpm"
	"github.com/bitwormhole/wpm/gen/main4wpm"
	"github.com/bitwormhole/wpm/gen/test4wpm"
	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/module-gorm-mysql/modules/mysql"
	"github.com/starter-go/module-gorm-sqlserver/modules/sqlserver"
	"github.com/starter-go/security/modules/security"
)

// Module  ...
func Module() application.Module {
	mb := wpm.NewMainModule()
	mb.Components(main4wpm.ExportComponents)

	mb.Depend(libgin.Module())
	mb.Depend(libgorm.Module())
	mb.Depend(security.Module())

	mb.Depend(mysql.Module())
	mb.Depend(sqlserver.Module())

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := wpm.NewTestModule()
	mb.Components(test4wpm.ExportComponents)
	return mb.Create()
}

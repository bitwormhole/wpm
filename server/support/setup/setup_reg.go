package setup

import (
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

type mySetupReg struct {
	context *Context
}

func (inst *mySetupReg) listAll() []*service.SetupRegistration {

	all := make([]*service.SetupRegistration, 0)

	all = append(all, inst.forInitMainRepository())
	all = append(all, inst.forDumpOlderSettings())
	all = append(all, inst.forImportPresetMediaFiles())
	all = append(all, inst.forImportPresetProjectTypes())
	all = append(all, inst.forImportPresetIntentTemplates())
	all = append(all, inst.forImportCommonExecutables())
	all = append(all, inst.forImportPresetSettings())
	all = append(all, inst.forImportOlderSettings())

	// all = append(all, inst.forExample())

	for i, item := range all {
		id := dxo.SetupID(i + 1)
		item.ID = id
		item.Prototype.ID = id
	}
	return all
}

func (inst *mySetupReg) forExample() *service.SetupRegistration {
	p := &dto.Setup{
		Name:        "example1",
		Title:       "Example1",
		Description: "This is a example-1",
		State:       dto.SetupStateInit,
	}
	h := &mySetupHanlders{context: inst.context}
	return &service.SetupRegistration{
		ID:        p.ID,
		Name:      p.Name,
		Prototype: p,
		Handler:   h.example,
	}
}

func (inst *mySetupReg) forImportPresetProjectTypes() *service.SetupRegistration {
	p := &dto.Setup{
		Name:        "import-preset-project-types",
		Title:       "导入预置项目类型",
		Description: "导入预置项目类型",
		State:       dto.SetupStateInit,
	}
	h := &mySetupHanlders{context: inst.context}
	return &service.SetupRegistration{
		ID:        p.ID,
		Name:      p.Name,
		Prototype: p,
		Handler:   h.doImportPresetProjectTypes,
	}
}

func (inst *mySetupReg) forImportPresetMediaFiles() *service.SetupRegistration {
	p := &dto.Setup{
		Name:        "import-preset-media-files",
		Title:       "导入预置媒体文件",
		Description: "导入预置媒体文件",
		State:       dto.SetupStateWant,
	}
	h := &mySetupHanlders{context: inst.context}
	return &service.SetupRegistration{
		ID:        p.ID,
		Name:      p.Name,
		Prototype: p,
		Handler:   h.doImportPresetMediaFiles,
	}
}

func (inst *mySetupReg) forInitMainRepository() *service.SetupRegistration {
	p := &dto.Setup{
		Name:        "init-main-repository",
		Title:       "初始化主仓库",
		Description: "初始化主仓库",
		State:       dto.SetupStateWant,
	}
	h := &mySetupHanlders{context: inst.context}
	return &service.SetupRegistration{
		ID:        p.ID,
		Name:      p.Name,
		Prototype: p,
		Handler:   h.doInitMainRepository,
	}
}

func (inst *mySetupReg) forImportPresetIntentTemplates() *service.SetupRegistration {
	p := &dto.Setup{
		Name:        "import-preset-intent-templates",
		Title:       "导入预置命令模板",
		Description: "导入预置命令模板",
		State:       dto.SetupStateWant,
	}
	h := &mySetupHanlders{context: inst.context}
	return &service.SetupRegistration{
		ID:        p.ID,
		Name:      p.Name,
		Prototype: p,
		Handler:   h.doImportPresetIntentTemplates,
	}
}

func (inst *mySetupReg) forImportCommonExecutables() *service.SetupRegistration {
	p := &dto.Setup{
		Name:        "import-common-executables",
		Title:       "导入常用程序",
		Description: "导入常用程序",
		State:       dto.SetupStateWant,
	}
	h := &mySetupHanlders{context: inst.context}
	return &service.SetupRegistration{
		ID:        p.ID,
		Name:      p.Name,
		Prototype: p,
		Handler:   h.doImportCommonExecutables,
	}
}

func (inst *mySetupReg) forImportPresetSettings() *service.SetupRegistration {
	p := &dto.Setup{
		Name:        "import-settings",
		Title:       "导入其它设置项",
		Description: "导入其它设置项",
		State:       dto.SetupStateWant,
	}
	h := &mySetupHanlders{context: inst.context}
	return &service.SetupRegistration{
		ID:        p.ID,
		Name:      p.Name,
		Prototype: p,
		Handler:   h.doImportPresetSettings,
	}
}

func (inst *mySetupReg) forDumpOlderSettings() *service.SetupRegistration {
	p := &dto.Setup{
		Name:        "dump-older-settings",
		Title:       "备份旧版本设置",
		Description: "备份旧版本设置",
		State:       dto.SetupStateWant,
	}
	h := &mySetupHanlders{context: inst.context}
	return &service.SetupRegistration{
		ID:        p.ID,
		Name:      p.Name,
		Prototype: p,
		Handler:   h.doDumpOlderData,
	}
}

func (inst *mySetupReg) forImportOlderSettings() *service.SetupRegistration {
	p := &dto.Setup{
		Name:        "import-older-settings",
		Title:       "导入旧版本设置",
		Description: "导入旧版本设置",
		State:       dto.SetupStateWant,
	}
	h := &mySetupHanlders{context: inst.context}
	return &service.SetupRegistration{
		ID:        p.ID,
		Name:      p.Name,
		Prototype: p,
		Handler:   h.doImportOlderData,
	}
}

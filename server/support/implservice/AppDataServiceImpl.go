package implservice

import (
	"bitwormhole.com/starter/afs"
	"bitwormhole.com/starter/afs/files"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
)

// AppDataServiceImpl ...
type AppDataServiceImpl struct {
	markup.Component `id:"AppDataService"`

	ProfileService service.ProfileService `inject:"#ProfileService"`

	baseAppDataDir afs.Path
}

func (inst *AppDataServiceImpl) _Impl() service.AppDataService {
	return inst
}

func (inst *AppDataServiceImpl) loadAppDataDir() afs.Path {
	profile, err := inst.ProfileService.GetProfile()
	if err != nil {
		panic(err)
	}
	home := files.FS().NewPath(profile.Home)
	return home.GetChild(".wpm")
}

func (inst *AppDataServiceImpl) forPath(child string) string {
	addir := inst.baseAppDataDir
	if addir == nil {
		addir = inst.loadAppDataDir()
		inst.baseAppDataDir = addir
	}
	path := addir.GetChild(child)
	return path.GetPath()
}

// GetAppDataDirectory 。。。
func (inst *AppDataServiceImpl) GetAppDataDirectory() string {
	return inst.forPath(".")
}

// GetSQLiteDBFile ...
func (inst *AppDataServiceImpl) GetSQLiteDBFile() string {
	return inst.forPath("wpm.db")
}

// GetMainRepositoryPath ...
func (inst *AppDataServiceImpl) GetMainRepositoryPath() string {
	return inst.forPath("main/.git")
}

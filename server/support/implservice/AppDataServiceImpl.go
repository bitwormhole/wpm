package implservice

import (
	"bitwormhole.com/starter/afs"
	"bitwormhole.com/starter/afs/files"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/service"
)

// 定义 app-data 文件名
const (
	AppDataFileDB  = "wpm.db"
	AppDataMainGit = "main/.git"
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

func (inst *AppDataServiceImpl) getAppDataDir() afs.Path {
	addir := inst.baseAppDataDir
	if addir == nil {
		addir = inst.loadAppDataDir()
		inst.baseAppDataDir = addir
	}
	return addir
}

func (inst *AppDataServiceImpl) forPath(child string) afs.Path {
	addir := inst.getAppDataDir()
	path := addir.GetChild(child)
	return path
}

func (inst *AppDataServiceImpl) forPathString(child string) string {
	return inst.forPath(child).GetPath()
}

// GetAppDataDirectory 。。。
func (inst *AppDataServiceImpl) GetAppDataDirectory() string {
	return inst.forPathString(".")
}

// GetSQLiteDBFile ...
func (inst *AppDataServiceImpl) GetSQLiteDBFile() string {
	return inst.forPathString(AppDataFileDB)
}

// GetMainRepositoryPath ...
func (inst *AppDataServiceImpl) GetMainRepositoryPath() string {
	return inst.forPathString(AppDataMainGit)
}

// Ready ...
func (inst *AppDataServiceImpl) Ready() bool {
	dir := inst.forPath(AppDataMainGit)
	return dir.Exists()
}

// Setup ...
func (inst *AppDataServiceImpl) Setup() error {
	dir := inst.forPath(AppDataMainGit)
	vlog.Info("setup main repository at ", dir.GetPath())
	opt := &afs.Options{}
	return dir.Mkdirs(opt)
}

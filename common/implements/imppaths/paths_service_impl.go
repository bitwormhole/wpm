package imppaths

import (
	"os"
	"os/user"

	"github.com/bitwormhole/wpm/common/classes/paths"
	"github.com/starter-go/afs"
)

// PathsServiceImpl 。。。
type PathsServiceImpl struct {

	//starter:component
	_as func(paths.Service) //starter:as("#")

	FS afs.FS //starter:inject("#")

	cached *paths.Locations
}

func (inst *PathsServiceImpl) _impl() paths.Service {
	return inst
}

// GetLocations ...
func (inst *PathsServiceImpl) GetLocations() paths.Locations {
	l := inst.cached
	if l == nil {
		l = inst.load()
		inst.cached = l
	}
	return *l
}

func (inst *PathsServiceImpl) load() *paths.Locations {

	l := new(paths.Locations)

	// 注意：这里的加载顺序有讲究 ...
	l.AppExeFile = inst.loadAppExeFile(l)
	l.UserHomeDir = inst.loadUserHomeDir(l)
	l.UserWPMDataDir = inst.loadUserWPMDataDir(l)
	l.MediaBucketPool = inst.loadMediaDir(l)

	return l
}

func (inst *PathsServiceImpl) loadUserHomeDir(l *paths.Locations) afs.Path {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	return inst.FS.NewPath(u.HomeDir)
}

func (inst *PathsServiceImpl) loadUserWPMDataDir(l *paths.Locations) afs.Path {
	return l.UserHomeDir.GetChild(".wpm")
}

func (inst *PathsServiceImpl) loadMediaDir(l *paths.Locations) afs.Path {
	return l.UserWPMDataDir.GetChild("files")
}

func (inst *PathsServiceImpl) loadAppExeFile(l *paths.Locations) afs.Path {
	arg0 := ""
	for i, ar := range os.Args {
		if i == 0 {
			arg0 = ar
			break
		}
	}
	if arg0 == "" {
		panic("exe file path is empty")
	}
	file := inst.FS.NewPath(arg0)
	if !file.IsFile() {
		panic("no exe file at path " + arg0)
	}
	return file
}

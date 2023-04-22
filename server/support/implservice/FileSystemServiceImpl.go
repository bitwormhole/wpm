package implservice

import (
	"bitwormhole.com/starter/afs"
	"bitwormhole.com/starter/afs/files"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
)

// FileSystemServiceImpl ...
type FileSystemServiceImpl struct {
	markup.Component `id:"FileSystemService"`

	fs afs.FS
}

func (inst *FileSystemServiceImpl) _Impl() service.FileSystemService {
	return inst
}

func (inst *FileSystemServiceImpl) loadFS() afs.FS {

	// xfs :=  files.FS()
	// xfs.NewPath( "" ) .GetIO().

	return files.FS()
}

// FS ...
func (inst *FileSystemServiceImpl) FS() afs.FS {
	fs := inst.fs
	if fs == nil {
		fs = inst.loadFS()
		inst.fs = fs
	}
	return fs
}

// Path ...
func (inst *FileSystemServiceImpl) Path(path string) afs.Path {
	return inst.FS().NewPath(path)
}

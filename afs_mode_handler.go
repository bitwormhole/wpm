package wpm

import (
	"io/fs"
	"os"

	"bitwormhole.com/starter/afs"
	"bitwormhole.com/starter/afs/files"
)

func InitAFSModeHandler() {
	h := &afsModeHandler{}
	files.FS().SetDefaultOptionsHandler(h.handle)
}

////////////////////////////////////////////////////////////////////////////////

type afsModeHandler struct {
}

func (inst *afsModeHandler) handle(path string, have *afs.Options, want afs.WantOption) *afs.Options {

	if have != nil {
		return have
	}

	opt := &afs.Options{
		Permission: fs.ModePerm,
		Flag:       0,
	}

	if inst.has(want, afs.WantToMakeDir) {
		opt.Mkdirs = true
	}

	if inst.has(want, afs.WantToReadFile) {
		opt.Flag |= os.O_RDONLY
	}

	if inst.has(want, afs.WantToWriteFile) {
		opt.Flag |= os.O_WRONLY
	}

	if inst.has(want, afs.WantToCreateFile) {
		opt.Flag |= os.O_CREATE
	}

	return opt
}

func (inst *afsModeHandler) has(a, b afs.WantOption) bool {
	return (a & b) != 0
}

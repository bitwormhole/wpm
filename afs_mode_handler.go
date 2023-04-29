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

func (inst *afsModeHandler) handle(path string, have, want *afs.Options) *afs.Options {

	if have != nil {
		return have
	}

	opt := &afs.Options{}

	if want != nil {
		if want.Write || want.Mkdirs || want.Create {
			opt.Flag |= os.O_WRONLY
		}
		if want.Read {
			opt.Flag |= os.O_RDONLY
		}
		if want.Create {
			opt.Permission = fs.ModePerm
			opt.Flag |= os.O_CREATE
		}
	}

	return opt
}

package utils

import (
	"testing"

	"bitwormhole.com/starter/afs/files"
)

func TestTarUnpack(t *testing.T) {

	tmp := files.FS().NewPath(t.TempDir())
	src := tmp.GetChild("src.tar.gz")
	dst := tmp.GetChild("dst.dir")

	if !src.IsFile() {
		t.Log("skip utils.TestTarUnpack(), no file ", src.GetPath())
		return
	}

	err := UnpackTar(src, dst)
	if err != nil {
		t.Error(err)
	}
}

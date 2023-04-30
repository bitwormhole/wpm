package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"

	"bitwormhole.com/starter/afs"
)

// Unzip ...
func Unzip(src afs.Path, dst afs.Path) error {
	u := &unzip{
		src: src,
		dst: dst,
	}
	return u.run()
}

////////////////////////////////////////////////////////////////////////////////

type unzip struct {
	src afs.Path
	dst afs.Path
}

func (inst *unzip) run() error {
	src := inst.src.GetPath()
	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		reader.Close()
	}()
	list := reader.File
	for _, f := range list {
		err = inst.handleItem(f)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *unzip) handleItem(f *zip.File) error {

	info := f.FileInfo()
	if info.IsDir() {
		return inst.mkdirs(f)
	}

	src, err := f.Open()
	if err != nil {
		return err
	}
	defer func() {
		src.Close()
	}()

	dst, err := inst.openDst(f)
	if err != nil {
		return err
	}
	defer func() {
		dst.Close()
	}()

	haveSize, err := io.Copy(dst, src)
	if err != nil {
		return err
	}
	wantSize := info.Size()
	if wantSize != haveSize {
		return fmt.Errorf("bad item size")
	}

	return nil
}

func (inst *unzip) openDst(f *zip.File) (io.WriteCloser, error) {
	opt := &afs.Options{
		Create: true, Write: true, File: true, Mkdirs: true,
		Flag:       os.O_CREATE | os.O_WRONLY,
		Permission: fs.ModePerm,
	}
	path := inst.dst.GetChild(f.Name)
	return path.GetIO().OpenWriter(opt)
}

func (inst *unzip) mkdirs(f *zip.File) error {
	opt := &afs.Options{
		Write: true, Directory: true,
	}
	path := inst.dst.GetChild(f.Name)
	return path.Mkdirs(opt)
}

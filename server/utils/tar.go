package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/starter/vlog"
)

// UnpackTar 解压tar包
func UnpackTar(src afs.Path, dst afs.Path) error {
	u := &unpacktar{
		src: src,
		dst: dst,
	}
	return u.run()
}

////////////////////////////////////////////////////////////////////////////////

type unpacktar struct {
	src afs.Path
	dst afs.Path
}

func (inst *unpacktar) run() error {
	return inst.unpackTarGz()
}

func (inst *unpacktar) unpackTarGz() error {

	// file reader
	r1, err := inst.src.GetIO().OpenReader(nil)
	if err != nil {
		return err
	}
	defer func() {
		r1.Close()
	}()

	// gzip reader
	r2, err := gzip.NewReader(r1)
	if err != nil {
		return err
	}
	defer func() {
		r2.Close()
	}()

	// tar reader
	r3 := tar.NewReader(r2)

	// loop
	for {
		h, err := r3.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		err = inst.handleItem(h, r3)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *unpacktar) handleItem(h *tar.Header, reader *tar.Reader) error {

	ok := false
	info := h.FileInfo()

	if info.IsDir() {
		inst.mkdirs(h)
		return nil
	}

	// dst
	dst, dstfile, err := inst.openDst(h)
	if err != nil {
		return err
	}
	defer func() {
		dst.Close()
		if ok {
			inst.setFileModeAndTime(dstfile, h)
		}
	}()

	src := reader
	wantSize := info.Size()
	haveSize, err := io.Copy(dst, src)
	if err != nil {
		return err
	}
	if wantSize != haveSize {
		return fmt.Errorf("bad item size")
	}

	ok = true
	return nil
}

func (inst *unpacktar) setFileModeAndTime(f afs.Path, h *tar.Header) {
	// todo ...
	// f.GetFS()
	vlog.Warn("todo: set file-mode & time, path=", f.GetPath())
}

func (inst *unpacktar) openDst(f *tar.Header) (io.WriteCloser, afs.Path, error) {
	mod := f.FileInfo().Mode()
	opt := &afs.Options{
		Mkdirs:     true,
		Permission: mod,
		Flag:       os.O_CREATE | os.O_WRONLY,
	}
	path := inst.dst.GetChild(f.Name)
	dir := path.GetParent()
	if !dir.Exists() {
		dir.Mkdirs(&afs.Options{Permission: os.ModePerm})
	}
	w, err := path.GetIO().OpenWriter(opt)
	if err != nil {
		return nil, path, err
	}
	return w, path, nil
}

func (inst *unpacktar) mkdirs(f *tar.Header) error {
	opt := &afs.Options{
		Write: true, Directory: true,
	}
	path := inst.dst.GetChild(f.Name)
	return path.Mkdirs(opt)
}

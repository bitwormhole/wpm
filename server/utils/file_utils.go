package utils

import (
	"io"

	"bitwormhole.com/starter/afs"
)

// CopyFile 复制文件内容
func CopyFile(src, dst afs.Path, writeOption *afs.Options) (int64, error) {

	in, err := src.GetIO().OpenReader(nil)
	if err != nil {
		return 0, err
	}
	defer in.Close()

	out, err := dst.GetIO().OpenWriter(writeOption)
	if err != nil {
		return 0, err
	}
	defer out.Close()

	return io.Copy(out, in)
}

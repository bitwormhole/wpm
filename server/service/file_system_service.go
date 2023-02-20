package service

import "bitwormhole.com/starter/afs"

// FileSystemService ...
type FileSystemService interface {
	FS() afs.FS

	Path(path string) afs.Path
}

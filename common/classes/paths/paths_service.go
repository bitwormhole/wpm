package paths

import "github.com/starter-go/afs"

// Locations ...
type Locations struct {
	UserHomeDir     afs.Path
	UserWPMDataDir  afs.Path
	MediaBucketPool afs.Path
	AppExeFile      afs.Path
}

// Service ...
type Service interface {
	GetLocations() Locations
}

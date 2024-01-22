package impmedia

import (
	"strings"

	"github.com/bitwormhole/wpm/common/classes/paths"
	"github.com/bitwormhole/wpm/server/classes/media"
	"github.com/starter-go/afs"
)

// MediaBucketPool 。。。
type MediaBucketPool struct {

	//starter:component
	_as func(media.BucketPool) //starter:as("#")

	Paths paths.Service //starter:inject("#")

}

func (inst *MediaBucketPool) _impl() media.BucketPool {
	return inst
}

// GetBucket ...
func (inst *MediaBucketPool) GetBucket(name string) media.Bucket {

	dir := inst.getMediaFolder()
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	if name == "" {
		name = "media"
	}

	b := new(bucketImpl)
	b.name = name
	b.dir = dir.GetChild(name)
	return b
}

func (inst *MediaBucketPool) getMediaFolder() afs.Path {
	loc := inst.Paths.GetLocations()
	return loc.MediaBucketPool
}

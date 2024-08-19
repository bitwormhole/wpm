package ibuckets

import (
	"github.com/bitwormhole/wpm"
	"github.com/bitwormhole/wpm/common/classes/buckets"
)

// MediaBucketPool 。。。
type MediaBucketPool struct {

	//starter:component
	_as func(buckets.BucketPool) //starter:as("#")

	Env wpm.Environment //starter:inject("#")

}

func (inst *MediaBucketPool) _impl() buckets.BucketPool {
	return inst
}

// GetBucket ...
func (inst *MediaBucketPool) GetBucket(name string) buckets.Bucket {

	dataDir := inst.Env.GetDataDir()
	bucketDir := dataDir.GetChild("buckets/" + name)

	b := new(bucketImpl)
	b.name = bucketDir.GetName()
	b.dir = bucketDir
	return b
}

package impmedia

import (
	"github.com/bitwormhole/wpm/server/classes/media"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/starter-go/afs"
)

// bucketImpl 。。。
type bucketImpl struct {
	name string
	dir  afs.Path
}

func (inst *bucketImpl) _impl() media.Bucket {
	return inst
}

func (inst *bucketImpl) GetName() string {
	return inst.name
}

func (inst *bucketImpl) GetPath() afs.Path {
	return inst.dir
}

func (inst *bucketImpl) GetMediaFile(m *dto.Media) afs.Path {
	i1 := 2
	name := m.SHA256SUM.String()
	count := len(name)
	if count <= i1 {
		return inst.dir.GetChild(name)
	}
	p1 := name[0:i1]
	p2 := name[i1:]
	return inst.dir.GetChild(p1 + "/" + p2)
}

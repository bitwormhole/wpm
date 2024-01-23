package buckets

import (
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/starter-go/afs"
)

// Bucket 表示一个存储媒体数据的桶（文件夹）
type Bucket interface {
	GetName() string

	GetPath() afs.Path

	GetMediaFile(m *dto.Media) afs.Path
}

// BucketPool 用于管理所有的媒体数据桶
type BucketPool interface {
	GetBucket(name string) Bucket
}

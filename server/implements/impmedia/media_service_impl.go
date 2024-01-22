package impmedia

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/bitwormhole/wpm/server/classes/media"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/starter-go/afs"
	"github.com/starter-go/base/lang"
)

// ServiceImpl ...
type ServiceImpl struct {

	//starter:component
	_as func(media.Service) //starter:as("#")

	Dao     media.DAO        //starter:inject("#")
	Buckets media.BucketPool //starter:inject("#")

	convertor media.Convertor
}

func (inst *ServiceImpl) _impl() media.Service {
	return inst
}

// Insert ...
func (inst *ServiceImpl) Insert(ctx context.Context, o1 *dto.Media) (*dto.Media, error) {
	o2 := inst.convertor.ConvertD2E(o1)
	o3, err := inst.Dao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}
	o4 := inst.convertor.ConvertE2D(o3)
	return o4, nil
}

// Update ...
func (inst *ServiceImpl) Update(ctx context.Context, id dxo.MediaID, o *dto.Media) (*dto.Media, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *ServiceImpl) Remove(ctx context.Context, id dxo.MediaID) error {
	return inst.Dao.Remove(nil, id)
}

// Find ...
func (inst *ServiceImpl) Find(ctx context.Context, id dxo.MediaID) (*dto.Media, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *ServiceImpl) List(ctx context.Context, q *media.Query) ([]*dto.Media, error) {
	list1, err := inst.Dao.List(nil, q)
	if err != nil {
		return nil, err
	}
	list2 := inst.convertor.ConvertListE2D(list1)
	return list2, nil
}

// ImportMediaStream ...
func (inst *ServiceImpl) ImportMediaStream(ctx context.Context, info *dto.Media, data io.Reader) (*dto.Media, error) {

	if ctx == nil || info == nil || data == nil {
		return nil, fmt.Errorf("param is nil")
	}

	bucket := inst.Buckets.GetBucket(info.Bucket)
	opt := inst.getWriteFileOptions()

	file1 := inst.getTempFile(info, bucket)
	file1.MakeParents(opt)
	err := inst.writeFile(file1, data)
	if err != nil {
		return nil, err
	}

	info.SHA256SUM = inst.computeSha256sum(file1)
	file2 := bucket.GetMediaFile(info)
	file2.MakeParents(opt)
	err = file1.MoveTo(file2, opt)
	if err != nil {
		return nil, err
	}

	info.Source = "file:///" + file2.GetPath()
	info.URL = inst.computeMediaURL(info)

	// return inst.Insert(ctx, info)
	return info, nil
}

func (inst *ServiceImpl) computeMediaURL(info *dto.Media) string {
	bucket := "media"
	filename := info.Name
	sum := info.SHA256SUM.String()
	parts := []string{bucket, "", "", sum, filename}
	const (
		ch0 = "/"
		ch1 = ";"
		sep = "\n"
	)
	ctype := info.ContentType
	ctype = strings.ReplaceAll(ctype, ch0, sep)
	ctype = strings.ReplaceAll(ctype, ch1, sep)
	typeNames := strings.Split(ctype, sep)
	for i, p := range typeNames {
		if i == 0 {
			parts[1] = p
		} else if i == 1 {
			parts[2] = p
		}
	}
	b := new(strings.Builder)
	for _, p := range parts {
		if p == "" {
			p = "null"
		}
		b.WriteString("/")
		b.WriteString(p)
	}
	return b.String()
}

func (inst *ServiceImpl) getTempFile(info *dto.Media, bucket media.Bucket) afs.Path {
	const (
		prefix = "import-media-"
		suffix = ".tmp"
	)
	dir := bucket.GetPath().GetChild("tmp")
	tmp, err := dir.GetFS().CreateTempFile(prefix, suffix, dir)
	if err != nil {
		now := lang.Now()
		str := strconv.FormatInt(now.Int(), 10)
		return dir.GetChild(prefix + str + suffix)
	}
	return tmp
}

func (inst *ServiceImpl) computeSha256sum(file afs.Path) lang.Hex {
	data, err := file.GetIO().ReadBinary(nil)
	if err != nil {
		return "0000"
	}
	sum := sha256.Sum256(data)
	return lang.HexFromBytes(sum[:])
}

func (inst *ServiceImpl) writeFile(file afs.Path, src io.Reader) error {
	opt := inst.getWriteFileOptions()
	dst, err := file.GetIO().OpenWriter(opt)
	if err != nil {
		return err
	}
	defer func() {
		dst.Close()
	}()
	_, err = io.Copy(dst, src)
	return err
}

func (inst *ServiceImpl) getWriteFileOptions() *afs.Options {
	ob := &afs.OptionsBuilder{}
	ob = ob.Create(true).Write(true).File(true).Mkdirs(true)
	return ob.Options()
}

// FindMediaFile ...
func (inst *ServiceImpl) FindMediaFile(ctx context.Context, want *dto.Media) (*dto.Media, afs.Path, error) {
	pool := inst.Buckets
	bucketName := want.Bucket
	bucket := pool.GetBucket(bucketName)
	file := bucket.GetMediaFile(want)
	if !file.IsFile() {
		return nil, nil, fmt.Errorf("no file")
	}
	return want, file, nil
}

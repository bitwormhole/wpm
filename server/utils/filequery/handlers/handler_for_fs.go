package handlers

import (
	"fmt"
	"net/url"
	"sort"
	"strings"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils/filequery"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

type FileSystemHandler struct {
	markup.Component `class:"filequery-handler-registry"`

	FS service.FileSystemService `inject:"#FileSystemService"`
}

func (inst *FileSystemHandler) _Impl() (filequery.HandlerRegistry, filequery.Handler) {
	return inst, inst
}

func (inst *FileSystemHandler) ListRegistrations() []*filequery.HandlerRegistration {
	r1 := &filequery.HandlerRegistration{
		Enabled: true,
		Handler: inst,
	}
	return []*filequery.HandlerRegistration{r1}
}

func (inst *FileSystemHandler) getfs() afs.FS {
	return inst.FS.FS()
}

func (inst *FileSystemHandler) Accept(q *vo.FileQuery) bool {
	url := q.BaseURL
	return strings.HasPrefix(url, "file:/")
}

func (inst *FileSystemHandler) queryRoots(q *vo.FileQuery) (*vo.FileQuery, error) {

	fs := inst.getfs()
	res := &vo.FileQuery{}
	roots := fs.ListRoots()
	items := make([]*dto.File, 0)

	for _, root := range roots {
		item := &dto.File{}
		inst.loadItem(root, item)
		// item.Name = ""
		items = append(items, item)
	}

	res.BaseURL = "file:///"
	res.Path = "/"
	res.Items = items
	res.Self.IsFile = false
	res.Self.Exists = true
	res.Self.IsDir = true
	return res, nil
}

func (inst *FileSystemHandler) Query(q *vo.FileQuery) (*vo.FileQuery, error) {

	u, err := url.Parse(q.BaseURL)
	if err != nil {
		return nil, err
	}
	pathStr := u.Path + "/" + q.Path
	fs := inst.getfs()
	path := fs.NewPath(pathStr)

	if path.GetPath() == "" {
		return inst.queryRoots(q)
	}

	res := &vo.FileQuery{}
	res.BaseURL = "file:///"
	res.Path = path.GetPath()

	err = inst.loadSelf(path, res)
	if err != nil {
		return nil, err
	}

	err = inst.loadItems(path, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (inst *FileSystemHandler) loadItems(path afs.Path, fq *vo.FileQuery) error {

	if fq == nil || path == nil {
		return fmt.Errorf("param is nil")
	}

	if !fq.Self.IsDir {
		// return fmt.Errorf("")
		return nil
	}

	dst := make([]*dto.File, 0)
	namelist := path.ListNames()
	sort.Strings(namelist)
	for _, name := range namelist {
		child := path.GetChild(name)
		item := &dto.File{}
		err := inst.loadItem(child, item)
		if err != nil {
			vlog.Warn(err)
		}
		item.Name = name
		dst = append(dst, item)
	}
	fq.Items = dst
	return nil
}

func (inst *FileSystemHandler) loadSelf(path afs.Path, fq *vo.FileQuery) error {
	item := &dto.File{}
	err := inst.loadItem(path, item)
	if err != nil {
		vlog.Warn(err)
	}
	item.Name = path.GetName()
	fq.Self = *item
	return nil
}

func (inst *FileSystemHandler) loadItem(path afs.Path, item *dto.File) error {

	info := path.GetInfo()

	item.Path = path.GetPath()
	item.Name = path.GetName()
	item.Size = info.Length()
	item.UpdatedAt = util.NewTime(info.UpdatedAt())

	item.IsDir = info.IsDirectory()
	item.IsFile = info.IsFile()
	item.Exists = info.Exists()

	if item.IsDir {
		item.Type = "directory"
	} else if item.IsFile {
		item.Type = "file"
	}

	return nil
}

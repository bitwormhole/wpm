package mediae

import (
	"context"
	"fmt"
	"strings"

	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
)

type myImportPresetsLoader struct {
	context context.Context
	parent  *MediaServiceImpl

	pathPrefixForRes string
	pathPrefixForWeb string
}

func (inst *myImportPresetsLoader) Load() error {

	err := inst.loadPrefixPaths()
	if err != nil {
		return err
	}

	list, err := inst.findPresets()
	if err != nil {
		return err
	}

	err = inst.insertPresets(list)
	if err != nil {
		return err
	}

	return nil
}

func (inst *myImportPresetsLoader) loadPrefixPaths() error {

	ser := inst.parent
	r := ser.ResPathPrefix
	w := ser.WebPathPrefix

	if r == "" {
		return fmt.Errorf("property [wpm.presets.res-path-prefix] is empty")
	}
	if w == "" {
		return fmt.Errorf("property [wpm.presets.web-path-prefix] is empty")
	}

	inst.pathPrefixForRes = r
	inst.pathPrefixForWeb = w
	return nil
}

func (inst *myImportPresetsLoader) findPresets() ([]*dto.Media, error) {
	dst := make([]*dto.Media, 0)
	res := inst.parent.AC.GetResources()
	all := res.Export(nil)
	for name, item := range all {
		if !item.IsFile() {
			continue
		}
		m, err := inst.makeMedia(name, item)
		if m != nil && err == nil {
			dst = append(dst, m)
		}
		if err != nil {
			vlog.Warn(err)
		}
	}
	return dst, nil
}

func (inst *myImportPresetsLoader) makeMedia(name string, r collection.Res) (*dto.Media, error) {

	prefixRes := inst.pathPrefixForRes
	// prefixWeb := inst.pathPrefixForWeb
	if !strings.HasPrefix(name, prefixRes) {
		return nil, nil // skip
	}
	// part2 := name[len(prefixRes):]
	// url := prefixWeb + part2

	bin, err := r.ReadBinary()
	if err != nil {
		return nil, err
	}

	size := len(bin)
	sum := utils.ComputeSHA256sum(bin)
	simpleName := inst.getSimpleFileName(name)
	ctype := inst.getContentType(name)
	source := "res:///" + name

	me := &dto.Media{
		Name:        simpleName,
		Label:       simpleName,
		URL:         "",
		Source:      source,
		FileSize:    int64(size),
		SHA256SUM:   sum,
		Bucket:      "preset",
		ContentType: ctype,
	}

	return me, nil
}

func (inst *myImportPresetsLoader) getContentType(name string) string {
	c := context.Background()
	ser := inst.parent.ContentTypeService
	t, err := ser.GetContentType(c, name)
	if err != nil {
		t = "application/octet-stream"
	}
	return t
}

func (inst *myImportPresetsLoader) getSimpleFileName(path string) string {
	path = strings.ReplaceAll(path, "\\", "/")
	i := strings.LastIndex(path, "/")
	if i < 0 {
		return path
	}
	return path[i+1:]
}

func (inst *myImportPresetsLoader) insertPresets(list []*dto.Media) error {
	cnt := 0
	for _, item := range list {
		err := inst.insertPreset(item)
		if err != nil {
			vlog.Warn(err)
		} else {
			cnt++
		}
	}
	vlog.Info("Has imported ", cnt, " media(s).")
	return nil
}

func (inst *myImportPresetsLoader) insertPreset(o1 *dto.Media) error {
	ctx := inst.context
	ser := inst.parent
	opt := &service.MediaOptions{
		FetchFromSource: true,
	}
	_, err := ser.Insert(ctx, o1, opt)
	return err
}

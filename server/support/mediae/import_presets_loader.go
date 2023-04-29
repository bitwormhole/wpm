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

	pathPrefixForRes []string
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
		return fmt.Errorf("property [wpm.presets.res-path-prefix-list] is empty")
	}
	if w == "" {
		return fmt.Errorf("property [wpm.presets.web-path-prefix] is empty")
	}

	inst.pathPrefixForRes = inst.parsePathPrefixForRes(r)
	inst.pathPrefixForWeb = w
	return nil
}

func (inst *myImportPresetsLoader) parsePathPrefixForRes(s string) []string {
	dst := make([]string, 0)
	src := strings.Split(s, ",")
	for _, item := range src {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		dst = append(dst, item)
	}
	return dst
}

func (inst *myImportPresetsLoader) findPresets() ([]*dto.Media, error) {
	dst := make([]*dto.Media, 0)
	res := inst.parent.AC.GetResources()
	all := res.Export(nil)
	for name, item := range all {
		if !inst.accept(name, item) {
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

func (inst *myImportPresetsLoader) accept(name string, item collection.Res) bool {

	prefixResList := inst.pathPrefixForRes
	prefixOK := false
	for _, prefix := range prefixResList {
		if strings.HasPrefix(name, prefix) {
			prefixOK = true
			break
		}
	}
	if !prefixOK {
		return false
	}

	if !item.IsFile() {
		return false
	}

	ignore := []string{".html", ".json", ".properties"}
	for _, suffix := range ignore {
		if strings.HasSuffix(name, suffix) {
			return false
		}
	}

	return true
}

func (inst *myImportPresetsLoader) makeMedia(name string, r collection.Res) (*dto.Media, error) {

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

package ifilequery

import (
	"context"
	"strings"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/server/classes/contenttypes"
)

type fileContentTypeInfoLoader struct {
	context      context.Context
	ctypeService contenttypes.Service

	all  map[string]*dto.ContentType
	used map[string]*dto.ContentType
}

func (inst *fileContentTypeInfoLoader) getUsedTable() map[string]*dto.ContentType {
	t := inst.used
	if t == nil {
		t = make(map[string]*dto.ContentType)
		inst.used = t
	}
	return t
}

func (inst *fileContentTypeInfoLoader) getAllTypeTable() map[string]*dto.ContentType {
	all := inst.all
	if all == nil {
		t, err := inst.loadTypeTable()
		if err != nil {
			return make(map[string]*dto.ContentType)
		}
		all = t
		inst.all = t
	}
	return all
}

func (inst *fileContentTypeInfoLoader) listUsedTypes() []*dto.ContentType {
	src := inst.getUsedTable()
	dst := make([]*dto.ContentType, 0)
	for _, ct := range src {
		dst = append(dst, ct)
	}
	return dst
}

func (inst *fileContentTypeInfoLoader) loadTypeTable() (map[string]*dto.ContentType, error) {
	ctx := inst.context
	dst := make(map[string]*dto.ContentType)
	q := &contenttypes.Query{All: true}
	list, err := inst.ctypeService.List(ctx, q)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		suffixList := item.Patterns.Array()
		for _, suffix := range suffixList {
			suffix = inst.normalizeSuffix(suffix)
			dst[suffix] = item
		}
	}

	return dst, nil
}

func (inst *fileContentTypeInfoLoader) loadForFile(file *dto.File) error {

	filename := file.Name
	suffix := inst.getFileNameSuffix(filename)
	all := inst.getAllTypeTable()
	used := inst.getUsedTable()

	t := all[suffix]
	if t == nil {
		file.Type = "file/" + suffix
		return nil
	}

	used[t.TypeName.String()] = t
	file.Type = t.TypeName.String()
	file.Icon = t.Icon
	return nil
}

func (inst *fileContentTypeInfoLoader) getFileNameSuffix(filename string) string {
	str, count := inst.getLastPartFromName(filename)
	if count <= 1 {
		return ""
	}
	return str
}

func (inst *fileContentTypeInfoLoader) normalizeSuffix(suffix string) string {
	str, _ := inst.getLastPartFromName(suffix)
	return str
}

func (inst *fileContentTypeInfoLoader) getLastPartFromName(name string) (string, int) {
	str := name
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	list := strings.Split(str, ".")
	count := len(list)
	return list[count-1], count
}

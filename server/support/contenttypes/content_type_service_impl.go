package contenttypes

import (
	"fmt"
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ContentTypeServiceImpl ...
type ContentTypeServiceImpl struct {
	markup.Component `id:"ContentTypeService"`

	AC            application.Context `inject:"context"`
	PropsFileName string              `inject:"${wpm.content-types.properties.filename}"`

	cache map[string]*dto.MIMEType
}

func (inst *ContentTypeServiceImpl) _Impl() service.ContentTypeService {
	return inst
}

func (inst *ContentTypeServiceImpl) loadCache() (map[string]*dto.MIMEType, error) {
	name := inst.PropsFileName
	text, err := inst.AC.GetResources().GetText(name)
	if err != nil {
		return nil, err
	}
	props, err := collection.ParseProperties(text, nil)
	if err != nil {
		return nil, err
	}
	return inst.loadTypes(props)
}

func (inst *ContentTypeServiceImpl) loadTypes(p collection.Properties) (map[string]*dto.MIMEType, error) {
	src := p.Export(nil)
	dst := make(map[string]*dto.MIMEType)
	for k, v := range src {
		t, err := inst.makeType(k, v)
		if t != nil && err == nil {
			dst[t.TypeName] = t
			dst[t.FileNameSuffix] = t
		}
		if err != nil {
			vlog.Warn(err)
		}
	}
	return dst, nil
}

func (inst *ContentTypeServiceImpl) makeType(k, v string) (*dto.MIMEType, error) {

	// [k:v] like 'content-types.suffix=type/subtype'
	const prefix = "content-types."
	k = strings.TrimSpace(k)
	v = strings.TrimSpace(v)
	k = strings.ToLower(k)
	v = strings.ToLower(v)

	if !strings.HasPrefix(k, prefix) {
		return nil, nil // skip
	}

	suffix := k[len(prefix)-1:]

	t := &dto.MIMEType{
		TypeName:       v,
		FileNameSuffix: suffix,
	}

	return t, nil
}

func (inst *ContentTypeServiceImpl) getCache() map[string]*dto.MIMEType {
	c := inst.cache
	if c != nil {
		return c
	}
	// load
	c, err := inst.loadCache()
	if err != nil {
		panic(err)
	}
	inst.cache = c
	return c
}

// GetContentType ...
func (inst *ContentTypeServiceImpl) GetContentType(name string) (string, error) {

	// pre-process suffix name
	i := strings.LastIndex(name, ".")
	if i < 0 {
		return "", fmt.Errorf("the filename is without a suffix like '.xxx'")
	}
	suffix := name[i:]
	suffix = strings.ToLower(suffix)
	suffix = strings.TrimSpace(suffix)

	// find in cache
	cache := inst.getCache()
	item := cache[suffix]
	if item == nil {
		// use default
		item = cache[".default"]
	}
	if item == nil {
		return "", fmt.Errorf("no type info for filename [%v]", name)
	}

	return item.TypeName, nil
}

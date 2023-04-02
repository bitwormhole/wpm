package contenttypes

import (
	"context"
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ProjectTypeImportServiceImpl ...
type ProjectTypeImportServiceImpl struct {
	markup.Component `id:"ProjectTypeImportService"`

	AC                 application.Context        `inject:"context"`
	ProjectTypeService service.ContentTypeService `inject:"#ContentTypeService"`
}

func (inst *ProjectTypeImportServiceImpl) _Impl() service.ProjectTypeImportService {
	return inst
}

// ImportTypesFromPreset ...
func (inst *ProjectTypeImportServiceImpl) ImportTypesFromPreset(c context.Context) error {
	all, err := inst.prepare(c)
	if err != nil {
		return err
	}
	return inst.insertAllTo(c, all)
}

func (inst *ProjectTypeImportServiceImpl) prepare(c context.Context) ([]*dto.ContentType, error) {

	const (
		path   = "res:///project-types.properties"
		prefix = "projecttype."
		suffix = ".type"
	)

	text, err := inst.AC.GetResources().GetText(path)
	if err != nil {
		return nil, err
	}

	props, err := collection.ParseProperties(text, nil)
	if err != nil {
		return nil, err
	}

	keys := inst.keys(props, prefix, suffix)
	dst := make([]*dto.ContentType, 0)
	for _, key := range keys {
		pt, err := inst.loadPT(props, prefix, key)
		if err == nil {
			dst = append(dst, pt)
		} else {
			vlog.Warn(err)
		}
	}
	return dst, nil
}

func (inst *ProjectTypeImportServiceImpl) loadPT(p collection.Properties, prefix string, key string) (*dto.ContentType, error) {

	k1 := prefix + key
	pt := &dto.ContentType{}
	getter := p.Getter()

	typeName := getter.GetString(k1+".type", "")
	typePattern := getter.GetString(k1+".patterns", "")

	pt.AsDir = getter.GetBool(k1+".as-dir", false)
	pt.AsFile = getter.GetBool(k1+".as-file", false)
	pt.Patterns = dxo.StringList(typePattern)
	pt.TypeName = dxo.ContentTypeName(typeName)
	pt.Label = getter.GetString(k1+".label", "")
	pt.Description = getter.GetString(k1+".description", "")
	pt.Priority = getter.GetInt(k1+".priority", 0)

	return pt, nil
}

func (inst *ProjectTypeImportServiceImpl) keys(p collection.Properties, prefix string, suffix string) []string {
	dst := make([]string, 0)
	table := p.Export(nil)
	for k := range table {
		if strings.HasPrefix(k, prefix) && strings.HasSuffix(k, suffix) {
			key := k[len(prefix) : len(k)-len(suffix)]
			dst = append(dst, key)
		}
	}
	return dst
}

func (inst *ProjectTypeImportServiceImpl) insertAllTo(c context.Context, all []*dto.ContentType) error {
	for _, item := range all {
		_, err := inst.insertOneTo(c, item)
		if err != nil {
			vlog.Warn(err)
		}
	}
	return nil
}

func (inst *ProjectTypeImportServiceImpl) insertOneTo(c context.Context, o *dto.ContentType) (*dto.ContentType, error) {
	ser := inst.ProjectTypeService
	return ser.Insert(c, o)
}

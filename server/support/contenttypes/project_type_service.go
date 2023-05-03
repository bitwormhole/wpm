package contenttypes

import (
	"context"
	"fmt"
	"strings"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ProjectTypeServiceImpl ...
type ProjectTypeServiceImpl struct {
	markup.Component `id:"ContentTypeService"`

	ProjectTypeDAO    dao.ContentTypeDAO        `inject:"#ContentTypeDAO"`
	FileSystemService service.FileSystemService `inject:"#FileSystemService"`

	internalContentTypes map[string]string // map[pattern] ctype
}

func (inst *ProjectTypeServiceImpl) _Impl() service.ContentTypeService {
	return inst
}

func (inst *ProjectTypeServiceImpl) dto2entity(o1 *dto.ContentType) (*entity.ContentType, error) {
	o2 := &entity.ContentType{}
	o2.ID = o1.ID
	o2.UUID = o1.UUID

	// todo ...
	o2.Name = o1.TypeName
	o2.Icon = o1.Icon
	o2.Patterns = o1.Patterns
	o2.Label = o1.Label
	o2.Description = o1.Description
	o2.AsDir = o1.AsDir
	o2.AsFile = o1.AsFile
	o2.AsProject = o1.AsProject
	o2.Priority = o1.Priority
	o2.Installation = o1.Installation

	o2.URN = dxo.NewContentTypeURN(o1.TypeName.String())

	return o2, nil
}

func (inst *ProjectTypeServiceImpl) entity2dto(o1 *entity.ContentType) (*dto.ContentType, error) {
	o2 := &dto.ContentType{}
	o2.ID = o1.ID
	o2.UUID = o1.UUID
	o2.URN = o1.URN
	o2.CreatedAt = util.NewTime(o1.CreatedAt)
	o2.UpdatedAt = util.NewTime(o1.UpdatedAt)

	// todo ...
	o2.TypeName = o1.Name
	o2.Icon = o1.Icon
	o2.Patterns = o1.Patterns
	o2.Label = o1.Label
	o2.Description = o1.Description
	o2.AsDir = o1.AsDir
	o2.AsFile = o1.AsFile
	o2.AsProject = o1.AsProject
	o2.Priority = o1.Priority
	o2.Installation = o1.Installation

	return o2, nil
}

// ListAll ...
func (inst *ProjectTypeServiceImpl) ListAll(ctx context.Context) ([]*dto.ContentType, error) {
	src, err := inst.ProjectTypeDAO.ListAll()
	if err != nil {
		return nil, err
	}
	dst := make([]*dto.ContentType, 0)
	for _, item1 := range src {
		item2, err := inst.entity2dto(item1)
		if err != nil {
			return nil, err
		}
		dst = append(dst, item2)
	}
	return dst, nil
}

// ListByPattern ...
func (inst *ProjectTypeServiceImpl) ListByPattern(ctx context.Context, pattern string) ([]*dto.ContentType, error) {
	src, err := inst.ProjectTypeDAO.ListByPattern(pattern)
	if err != nil {
		return nil, err
	}
	dst := make([]*dto.ContentType, 0)
	for _, item1 := range src {
		item2, err := inst.entity2dto(item1)
		if err != nil {
			return nil, err
		}
		dst = append(dst, item2)
	}
	return dst, nil
}

// ListByPatterns ...
func (inst *ProjectTypeServiceImpl) ListByPatterns(ctx context.Context, patterns []string) ([]*dto.ContentType, error) {
	// 简单实现：临时用于测试
	dst := make([]*dto.ContentType, 0)
	for _, p := range patterns {
		src, err := inst.ListByPattern(ctx, p)
		if err != nil {
			vlog.Warn(err)
			continue
		}
		if src == nil {
			continue
		}
		dst = append(dst, src...)
	}
	return dst, nil
}

// Find ...
func (inst *ProjectTypeServiceImpl) Find(ctx context.Context, id dxo.ContentTypeID) (*dto.ContentType, error) {
	o1, err := inst.ProjectTypeDAO.Find(id)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o1)
}

// Insert ...
func (inst *ProjectTypeServiceImpl) Insert(ctx context.Context, o1 *dto.ContentType) (*dto.ContentType, error) {
	o2, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}

	// prepare for insert
	o2.URN = dxo.NewContentTypeURN(o2.Name.String())

	o3, err := inst.ProjectTypeDAO.Insert(o2)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o3)
}

// Update ...
func (inst *ProjectTypeServiceImpl) Update(ctx context.Context, id dxo.ContentTypeID, o1 *dto.ContentType) (*dto.ContentType, error) {
	o2, err := inst.dto2entity(o1)
	if err != nil {
		return nil, err
	}
	o3, err := inst.ProjectTypeDAO.Update(id, o2)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o3)
}

// Remove ...
func (inst *ProjectTypeServiceImpl) Remove(ctx context.Context, id dxo.ContentTypeID) error {
	return inst.ProjectTypeDAO.Remove(id)
}

// LocateProject ...
func (inst *ProjectTypeServiceImpl) LocateProject(ctx context.Context, o *dto.Project, path string) error {
	locator := &myProjectLocatorWithTypes{
		parent:  inst,
		path:    path,
		project: o,
		context: ctx,
	}
	return locator.Locate()
}

// GetContentType 根据文件名（或路径）查询对应的mime类型
func (inst *ProjectTypeServiceImpl) GetContentType(ctx context.Context, name string) (string, error) {

	pattern := inst.getPatternByName(name)
	ct, err := inst.findContentType(ctx, pattern)
	if err == nil && ct != "" {
		return ct, nil
	}

	// 再查一下内部类型，看看有没有
	table := inst.getInternalTypes()
	ct = table[pattern]
	if ct == "" {
		return "", fmt.Errorf("no content-type info for file name:" + name)
	}

	return ct, nil
}

func (inst *ProjectTypeServiceImpl) getPatternByName(name string) string {
	// pre-process suffix name
	i := strings.LastIndex(name, ".")
	if i < 0 {
		// return "", fmt.Errorf("the filename is without a suffix like '.xxx'")
		return ""
	}
	suffix := name[i:]
	suffix = strings.ToLower(suffix)
	suffix = strings.TrimSpace(suffix)
	pattern := "*" + suffix
	return pattern
}

func (inst *ProjectTypeServiceImpl) findContentType(ctx context.Context, pattern string) (string, error) {
	list, err := inst.ListByPattern(ctx, pattern)
	if err != nil {
		return "", err
	}
	if len(list) < 1 {
		return "", fmt.Errorf("no content-type info for name:" + pattern)
	}
	item := list[0]
	return item.TypeName.String(), nil
}

func (inst *ProjectTypeServiceImpl) getInternalTypes() map[string]string {
	table := inst.internalContentTypes
	if table == nil {
		table = inst.loadInternalTypes()
		inst.internalContentTypes = table
	}
	return table
}

func (inst *ProjectTypeServiceImpl) loadInternalTypes() map[string]string {
	table := inst.internalContentTypes
	table = make(map[string]string)

	table["*.txt"] = "text/plain"
	table["*.html"] = "text/html"

	table["*.gif"] = "image/gif"
	table["*.png"] = "image/png"
	table["*.jpg"] = "image/jpeg"
	table["*.svg"] = "image/svg+xml"

	return table
}

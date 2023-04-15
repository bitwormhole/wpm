package jsonbuffer

import (
	"context"
	"fmt"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// ImpJSONBufferService ...
type ImpJSONBufferService struct {
	markup.Component `id:"JSONBufferService"`

	view *vo.Online
}

func (inst *ImpJSONBufferService) _Impl() service.JSONBufferService {
	return inst
}

func (inst *ImpJSONBufferService) makeNewVO() *vo.Online {
	o := &vo.Online{}

	o.ContentTypes = make([]*dto.ContentType, 0)
	o.Executables = make([]*dto.Executable, 0)
	o.IntentTemplates = make([]*dto.IntentTemplate, 0)
	o.Mediae = make([]*dto.Media, 0)
	o.Packages = make([]*dto.SoftwarePackage, 0)
	o.Sources = make([]*dto.Namespace, 0)

	return o
}

// Get ...
func (inst *ImpJSONBufferService) Get(ctx context.Context) (*vo.Online, error) {
	o := inst.view
	if o == nil {
		o = inst.makeNewVO()
		inst.view = o
	}
	return o, nil
}

// Reset ...
func (inst *ImpJSONBufferService) Reset(ctx context.Context) error {
	inst.view = inst.makeNewVO()
	return nil
}

// Put ...
func (inst *ImpJSONBufferService) Put(ctx context.Context, src *vo.Online) error {

	if src == nil {
		return fmt.Errorf("param:src is nil")
	}

	dst, _ := inst.Get(ctx)

	dst.Executables = append(dst.Executables, src.Executables...)
	dst.IntentTemplates = append(dst.IntentTemplates, src.IntentTemplates...)
	dst.ContentTypes = append(dst.ContentTypes, src.ContentTypes...)
	dst.Mediae = append(dst.Mediae, src.Mediae...)
	dst.Sources = append(dst.Sources, src.Sources...)
	dst.Packages = append(dst.Packages, src.Packages...)

	return nil
}

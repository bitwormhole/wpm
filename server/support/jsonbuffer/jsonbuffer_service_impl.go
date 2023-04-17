package jsonbuffer

import (
	"context"
	"fmt"
	"runtime"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
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

	pack := &dto.SoftwarePackage{
		Namespace:   "default",
		Name:        "todo",
		OS:          runtime.GOOS,
		Arch:        runtime.GOARCH,
		Title:       "todo",
		Description: "todo",
		Icon:        "https://todo",
		Revision:    1,
		Version:     "0.0.0",
	}
	o.Packages = append(o.Packages, pack)

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

	// append
	dst.Executables = append(dst.Executables, src.Executables...)
	dst.IntentTemplates = append(dst.IntentTemplates, src.IntentTemplates...)
	dst.ContentTypes = append(dst.ContentTypes, src.ContentTypes...)
	dst.Mediae = append(dst.Mediae, src.Mediae...)
	dst.Sources = append(dst.Sources, src.Sources...)
	dst.Packages = append(dst.Packages, src.Packages...)

	// sort
	inst.sortContentTypes(dst)
	inst.sortExecutables(dst)
	inst.sortIntentTemplates(dst)
	inst.sortMediae(dst)
	inst.sortPackages(dst)
	inst.sortSources(dst)

	return nil
}

func (inst *ImpJSONBufferService) sortContentTypes(o *vo.Online) {

	// 排序
	src := o.ContentTypes
	sa := &utils.SortAdapter{Quietly: true}
	sa.OnLen = func() int {
		return len(src)
	}
	sa.OnLess = func(i1, i2 int) bool {
		t1 := src[i1].UpdatedAt
		t2 := src[i2].UpdatedAt
		return t1 > t2
	}
	sa.OnSwap = func(i1, i2 int) {
		o1 := src[i1]
		o2 := src[i2]
		src[i1] = o2
		src[i2] = o1
	}
	sa.Sort()

	// 排重
	dst := make([]*dto.ContentType, 0)
	set := make(map[dxo.UUID]bool)
	for _, item := range src {
		uuid := item.UUID
		if set[uuid] {
			continue
		}
		set[uuid] = true
		dst = append(dst, item)
	}

	o.ContentTypes = dst
}

func (inst *ImpJSONBufferService) sortIntentTemplates(o *vo.Online) {

	// 排序
	src := o.IntentTemplates
	sa := &utils.SortAdapter{Quietly: true}
	sa.OnLen = func() int {
		return len(src)
	}
	sa.OnLess = func(i1, i2 int) bool {
		t1 := src[i1].UpdatedAt
		t2 := src[i2].UpdatedAt
		return t1 > t2
	}
	sa.OnSwap = func(i1, i2 int) {
		o1 := src[i1]
		o2 := src[i2]
		src[i1] = o2
		src[i2] = o1
	}
	sa.Sort()

	// 排重
	dst := make([]*dto.IntentTemplate, 0)
	set := make(map[dxo.UUID]bool)
	for _, item := range src {
		uuid := item.UUID
		if set[uuid] {
			continue
		}
		set[uuid] = true
		dst = append(dst, item)
	}

	o.IntentTemplates = dst
}

func (inst *ImpJSONBufferService) sortMediae(o *vo.Online) {

	// 排序
	src := o.Mediae
	sa := &utils.SortAdapter{Quietly: true}
	sa.OnLen = func() int {
		return len(src)
	}
	sa.OnLess = func(i1, i2 int) bool {
		t1 := src[i1].UpdatedAt
		t2 := src[i2].UpdatedAt
		return t1 > t2
	}
	sa.OnSwap = func(i1, i2 int) {
		o1 := src[i1]
		o2 := src[i2]
		src[i1] = o2
		src[i2] = o1
	}
	sa.Sort()

	// 排重
	dst := make([]*dto.Media, 0)
	set := make(map[dxo.UUID]bool)
	for _, item := range src {
		uuid := item.UUID
		if set[uuid] {
			continue
		}
		set[uuid] = true
		dst = append(dst, item)
	}

	o.Mediae = dst
}

func (inst *ImpJSONBufferService) sortExecutables(o *vo.Online) {

	// 排序
	src := o.Executables
	sa := &utils.SortAdapter{Quietly: true}
	sa.OnLen = func() int {
		return len(src)
	}
	sa.OnLess = func(i1, i2 int) bool {
		t1 := src[i1].UpdatedAt
		t2 := src[i2].UpdatedAt
		return t1 > t2
	}
	sa.OnSwap = func(i1, i2 int) {
		o1 := src[i1]
		o2 := src[i2]
		src[i1] = o2
		src[i2] = o1
	}
	sa.Sort()

	// 排重
	dst := make([]*dto.Executable, 0)
	set := make(map[dxo.UUID]bool)
	for _, item := range src {
		uuid := item.UUID
		if set[uuid] {
			continue
		}
		set[uuid] = true
		dst = append(dst, item)
	}

	o.Executables = dst
}

func (inst *ImpJSONBufferService) sortPackages(o *vo.Online) {

	// 排序
	src := o.Packages
	sa := &utils.SortAdapter{Quietly: true}
	sa.OnLen = func() int {
		return len(src)
	}
	sa.OnLess = func(i1, i2 int) bool {
		t1 := src[i1].UpdatedAt
		t2 := src[i2].UpdatedAt
		return t1 > t2
	}
	sa.OnSwap = func(i1, i2 int) {
		o1 := src[i1]
		o2 := src[i2]
		src[i1] = o2
		src[i2] = o1
	}
	sa.Sort()

	// 排重
	dst := make([]*dto.SoftwarePackage, 0)
	set := make(map[dxo.UUID]bool)
	for _, item := range src {
		uuid := item.UUID
		if set[uuid] {
			continue
		}
		set[uuid] = true
		dst = append(dst, item)
	}

	o.Packages = dst
}

func (inst *ImpJSONBufferService) sortSources(o *vo.Online) {

	// 排序
	src := o.Sources
	sa := &utils.SortAdapter{Quietly: true}
	sa.OnLen = func() int {
		return len(src)
	}
	sa.OnLess = func(i1, i2 int) bool {
		t1 := src[i1].UpdatedAt
		t2 := src[i2].UpdatedAt
		return t1 > t2
	}
	sa.OnSwap = func(i1, i2 int) {
		o1 := src[i1]
		o2 := src[i2]
		src[i1] = o2
		src[i2] = o1
	}
	sa.Sort()

	// 排重
	dst := make([]*dto.Namespace, 0)
	set := make(map[dxo.UUID]bool)
	for _, item := range src {
		uuid := item.UUID
		if set[uuid] {
			continue
		}
		set[uuid] = true
		dst = append(dst, item)
	}

	o.Sources = dst
}

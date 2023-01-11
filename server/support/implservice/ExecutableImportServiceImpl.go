package implservice

import (
	"context"
	"errors"
	"strings"

	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// ExecutableImportServiceImpl ...
type ExecutableImportServiceImpl struct {
	markup.Component `id:"ExecutableImportService"`

	ExecutableService service.ExecutableService `inject:"#ExecutableService"`
}

func (inst *ExecutableImportServiceImpl) _Impl() service.ExecutableImportService {
	return inst
}

// Save ...
func (inst *ExecutableImportServiceImpl) Save(ctx context.Context, o1 *vo.ExecutableImport) (*vo.ExecutableImport, error) {

	ser := inst.ExecutableService
	items := o1.Executables
	// errlist := make([]*dto.Executable, 0)
	var err2 error

	for _, item := range items {
		_, err := ser.Insert(ctx, item)
		if err != nil {
			// errlist = append(errlist, item)
			err2 = err
		}
	}

	if err2 != nil {
		return nil, err2
	}

	o2 := &vo.ExecutableImport{}
	return o2, nil
}

// Locate ...
func (inst *ExecutableImportServiceImpl) Locate(ctx context.Context, o1 *vo.ExecutableImport) (*vo.ExecutableImport, error) {
	src := o1.Executables
	dst := make([]*dto.Executable, 0)
	fsys := fs.Default()
	for _, item := range src {
		file := fsys.GetPath(item.Path)
		info, err := inst.getExeInfo(ctx, file)
		if err != nil {
			info = item
			info.Remark = err.Error()
			info.Ready = false
		}
		dst = append(dst, info)
	}
	o2 := &vo.ExecutableImport{}
	o2.Executables = dst
	return o2, nil
}

func (inst *ExecutableImportServiceImpl) getExeInfo(ctx context.Context, file fs.Path) (*dto.Executable, error) {

	if !file.IsFile() {
		return nil, errors.New("")
	}

	sum, err := utils.ComputeFileSHA256sum(file)
	if err != nil {
		return nil, err
	}

	// get name & title
	name := file.Name()
	title := strings.ToLower(name)
	const suffix = ".exe"
	if strings.HasSuffix(title, suffix) {
		title = title[0 : len(title)-len(suffix)]
	}

	// make dto
	o := &dto.Executable{}
	o.Size = file.Size()
	o.SHA256SUM = sum
	o.Title = title
	o.Name = name
	o.Path = file.Path()
	o.Ready = true
	o.Remark = ""
	o.ID = 0

	return o, nil
}

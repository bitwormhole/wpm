package executables

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ExecutableServiceImpl ...
type ExecutableServiceImpl struct {
	markup.Component `id:"ExecutableService"`

	// IconService       service.AppIconService    `inject:"#AppIconService"` // 【已弃用】

	ExecutableDAO     dao.ExecutableDAO         `inject:"#ExecutableDAO"`
	FileSystemService service.FileSystemService `inject:"#FileSystemService"`
	LocationService   service.LocationService   `inject:"#LocationService"`
}

func (inst *ExecutableServiceImpl) _Impl() service.ExecutableService {
	return inst
}

func (inst *ExecutableServiceImpl) normalizeOptions(opt *service.ExecutableOptions) *service.ExecutableOptions {
	if opt == nil {
		opt = &service.ExecutableOptions{}
	}
	return opt
}

func (inst *ExecutableServiceImpl) prepareLocation(c context.Context, o1 *entity.Executable) error {
	path := o1.Path
	location := &dto.Location{
		Path:   path,
		Class:  dxo.LocationExecutable,
		AsFile: true,
		AsDir:  true,
	}
	location, err := inst.LocationService.InsertOrFetch(c, location, nil)
	if err != nil {
		return err
	}
	// o1.Path = location.Path
	o1.Location = location.ID
	o1.Class = location.Class
	return nil
}

func (inst *ExecutableServiceImpl) dto2entity(c context.Context, o1 *dto.Executable, opt *service.ExecutableOptions) (*entity.Executable, error) {

	opt = inst.normalizeOptions(opt)

	o2 := &entity.Executable{}
	o2.ID = o1.ID
	o2.URN = o1.URN
	o2.Referer = o1.Referer

	o2.Name = o1.Name
	o2.Aliases = o1.Aliases
	o2.Namespace = o1.Namespace
	o2.Title = o1.Title
	o2.Size = o1.Size
	o2.SHA256SUM = o1.SHA256SUM
	o2.Description = o1.Description
	o2.IconURL = o1.IconURL
	o2.OpenWithPriority = o1.OpenWithPriority

	o2.Path = o1.Path
	o2.Location = o1.Location
	o2.Class = o1.Class

	// compute
	err := inst.computeEntity1(o2, opt)
	if err != nil {
		return nil, err
	}
	err = inst.computeEntity2(o2, opt)
	if err != nil {
		return nil, err
	}

	return o2, nil
}

// computeEntity1: 根据文件路径计算各个字段
func (inst *ExecutableServiceImpl) computeEntity1(o2 *entity.Executable, opt *service.ExecutableOptions) error {

	opt = inst.normalizeOptions(opt)
	if opt.SkipFileChecking {
		return nil
	}

	file := inst.FileSystemService.Path(o2.Path)
	name := file.GetName()

	sum, err := utils.ComputeFileSHA256sumAFS(file)
	if err != nil {
		return err
	}

	o2.Name = name
	o2.Size = file.GetInfo().Length()
	o2.SHA256SUM = sum
	return nil
}

// computeEntity2: 根据名称计算各个字段
func (inst *ExecutableServiceImpl) computeEntity2(o2 *entity.Executable, opt *service.ExecutableOptions) error {

	// opt = inst.normalizeOptions(opt)

	ns := o2.Namespace
	name := o2.Name
	normalName := inst.normalizeExeName(name)

	aliases := o2.Aliases.Array()
	aliases = append(aliases, name)
	aliases = append(aliases, normalName)
	aliases = inst.paiChong(aliases)

	if ns == "" {
		ns = "default" // the default NS
	}

	o2.Namespace = ns
	o2.Name = normalName
	o2.URN = dxo.NewExecutableURN(ns + "#" + normalName)
	o2.Aliases = dxo.NewStringList(aliases)

	return nil
}

func (inst *ExecutableServiceImpl) normalizeExeName(name string) string {
	const suffixExe = ".exe"
	name = strings.ToLower(name)
	name = strings.TrimSpace(name)
	if strings.HasSuffix(name, suffixExe) {
		return name[0 : len(name)-len(suffixExe)]
	}
	return name
}

// 排重
func (inst *ExecutableServiceImpl) paiChong(src []string) []string {
	dst := make([]string, 0)
	table := make(map[string]string)
	for _, str := range src {
		str = strings.TrimSpace(str)
		if str != "" {
			table[str] = str
		}
	}
	for _, str := range table {
		dst = append(dst, str)
	}
	sort.Strings(dst)
	return dst
}

func (inst *ExecutableServiceImpl) entity2dto(o1 *entity.Executable) (*dto.Executable, error) {

	o2 := &dto.Executable{}
	o2.ID = o1.ID
	o2.UUID = o1.UUID
	o2.CreatedAt = util.NewTime(o1.CreatedAt)
	o2.UpdatedAt = util.NewTime(o1.UpdatedAt)
	o2.URN = o1.URN
	o2.Referer = o1.Referer

	o2.Name = o1.Name
	o2.Aliases = o1.Aliases
	o2.Namespace = o1.Namespace
	o2.Title = o1.Title
	o2.Size = o1.Size
	o2.SHA256SUM = o1.SHA256SUM
	o2.Description = o1.Description
	o2.IconURL = o1.IconURL
	o2.OpenWithPriority = o1.OpenWithPriority

	o2.Path = o1.Path
	o2.Location = o1.Location
	o2.Class = o1.Class

	// todo ...
	o2.State = inst.checkExeFileState(o1)

	// inst.IconService.FillWithIconURL(o2)

	return o2, nil
}

func (inst *ExecutableServiceImpl) checkExeFileState(o1 *entity.Executable) dxo.FileState {
	path := inst.FileSystemService.Path(o1.Path)
	if path.IsFile() {
		return dxo.FileStateReady
	} else if path.IsDirectory() {
		return dxo.FileStateUnknowError
	}
	return dxo.FileStateOffline
}

func (inst *ExecutableServiceImpl) checkBeforeInsert(ctx context.Context, o *dto.Executable, opt *service.ExecutableOptions) error {

	opt = inst.normalizeOptions(opt)
	if opt.SkipFileChecking {
		return nil
	}

	path := fs.Default().GetPath(o.Path)
	if path == nil {
		return fmt.Errorf("bad file path: [%v]", o.Path)
	}
	if !path.IsFile() {
		return errors.New("the file is not a executable, path=" + o.Path)
	}
	old, _ := inst.FindByPath(ctx, o.Path, opt)
	if old != nil {
		return errors.New("the executable is exists, path=" + o.Path)
	}
	return nil
}

// ListAll ...
func (inst *ExecutableServiceImpl) ListAll(ctx context.Context, opt *service.ExecutableOptions) ([]*dto.Executable, error) {
	src, err := inst.ExecutableDAO.ListAll()
	if err != nil {
		return nil, err
	}
	dst := make([]*dto.Executable, 0)
	for _, item1 := range src {
		item2, err := inst.entity2dto(item1)
		if err != nil {
			return nil, err
		}
		dst = append(dst, item2)
	}
	return dst, nil
}

// Find ...
func (inst *ExecutableServiceImpl) Find(ctx context.Context, id dxo.ExecutableID, opt *service.ExecutableOptions) (*dto.Executable, error) {
	o1, err := inst.ExecutableDAO.Find(id)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o1)
}

// Insert ...
func (inst *ExecutableServiceImpl) Insert(ctx context.Context, o1 *dto.Executable, opt *service.ExecutableOptions) (*dto.Executable, error) {
	err := inst.checkBeforeInsert(ctx, o1, opt)
	if err != nil {
		return nil, err
	}
	e1, err := inst.dto2entity(ctx, o1, opt)
	if err != nil {
		return nil, err
	}
	err = inst.prepareLocation(ctx, e1)
	if err != nil {
		return nil, err
	}
	e2, err := inst.ExecutableDAO.Insert(e1)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(e2)
}

// Update ...
func (inst *ExecutableServiceImpl) Update(ctx context.Context, id dxo.ExecutableID, o1 *dto.Executable, opt *service.ExecutableOptions) (*dto.Executable, error) {
	o2, err := inst.dto2entity(ctx, o1, opt)
	if err != nil {
		return nil, err
	}
	err = inst.prepareLocation(ctx, o2)
	if err != nil {
		return nil, err
	}
	o3, err := inst.ExecutableDAO.Update(id, o2)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(o3)
}

// Remove ...
func (inst *ExecutableServiceImpl) Remove(ctx context.Context, id dxo.ExecutableID, opt *service.ExecutableOptions) error {
	return inst.ExecutableDAO.Remove(id)
}

// FindByPath ...
func (inst *ExecutableServiceImpl) FindByPath(ctx context.Context, path string, opt *service.ExecutableOptions) (*dto.Executable, error) {
	e1, err := inst.ExecutableDAO.FindByPath(path)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(e1)
}

// FindByName ...
func (inst *ExecutableServiceImpl) FindByName(ctx context.Context, name string, opt *service.ExecutableOptions) (*dto.Executable, error) {
	e1, err := inst.ExecutableDAO.FindByName(name)
	if err != nil {
		return nil, err
	}
	return inst.entity2dto(e1)
}

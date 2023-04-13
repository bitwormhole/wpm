package appruntime

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/backup"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
)

// ImpAppRuntimeService ...
type ImpAppRuntimeService struct {
	markup.Component `id:"AppRuntimeService" class:"life"`

	FileSystemService service.FileSystemService `inject:"#FileSystemService"`
	AppDataService    service.AppDataService    `inject:"#AppDataService"`
	MediaService      service.MediaService      `inject:"#MediaService"`

	EnableBackupSelf bool `inject:"${wpm.options.backup-this-exe}"`

	info *entity.Executable // the cached runtime info
}

func (inst *ImpAppRuntimeService) _Impl() (service.AppRuntimeService, application.LifeRegistry) {
	return inst, inst
}

// GetLifeRegistration ...
func (inst *ImpAppRuntimeService) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnInit: inst.onInit,
	}
}

func (inst *ImpAppRuntimeService) onInit() error {
	inst.GetAppHash()
	inst.logSelf()
	inst.backupSelf()
	return nil
}

// GetAppHash 取当前 exe 文件的 sha-256
func (inst *ImpAppRuntimeService) GetAppHash() (util.Hex, error) {
	i, err := inst.getAppRuntimeInfo()
	if err != nil {
		return "", err
	}
	return i.SHA256SUM, nil
}

func (inst *ImpAppRuntimeService) getAppRuntimeInfo() (*entity.Executable, error) {
	info := inst.info
	if info != nil {
		return info, nil
	}
	info, err := inst.loadAppRuntimeInfo()
	if err != nil {
		return nil, err
	}
	inst.info = info
	return info, nil
}

func (inst *ImpAppRuntimeService) loadAppRuntimeInfo() (*entity.Executable, error) {

	info := &entity.Executable{}
	now := time.Now()
	path := os.Args[0]

	// 重新计算
	file := inst.FileSystemService.Path(path)
	hex, err := utils.ComputeFileSHA256sumAFS(file)
	if err != nil {
		return nil, err
	}

	mod := service.GetAppModule()
	ver := mod.GetVersion()
	rev := mod.GetRevision()

	info.Path = path
	info.SHA256SUM = hex
	info.CreatedAt = now
	info.UpdatedAt = now
	info.Name = file.GetName()
	info.Size = file.GetInfo().Length()
	info.Title = mod.GetName()
	info.Description = "wpm " + ver + " r" + strconv.Itoa(rev)

	vlog.Info("[exe sha256sum:", hex, " path:", path, "]")
	return info, nil
}

func (inst *ImpAppRuntimeService) logSelf() error {
	mod := service.GetAppModule()
	o1 := &backup.StartupVO{}
	o2, err := inst.getAppRuntimeInfo()
	if err != nil {
		return err
	}
	inst.loadLogs(o1)
	if o1.Versions == nil {
		o1.Versions = make(map[string]*entity.Executable)
	}
	o1.WPM.Name = mod.GetName()
	o1.WPM.Version = mod.GetVersion()
	o1.WPM.Revision = mod.GetRevision()
	o1.WPM.HexName = o2.SHA256SUM.String()
	o1.Versions[o2.SHA256SUM.String()] = o2
	// o1.Logs = append(o1.Logs, o2)
	return inst.saveLogs(o1)
}

func (inst *ImpAppRuntimeService) loadLogs(o *backup.StartupVO) error {
	file := inst.getLogFile()
	data, err := file.GetIO().ReadBinary(nil)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, o)
}

func (inst *ImpAppRuntimeService) saveLogs(o *backup.StartupVO) error {
	data, err := json.Marshal(o)
	if err != nil {
		return err
	}
	file := inst.getLogFile()
	opt := &afs.Options{Mkdirs: true, Create: true}
	return file.GetIO().WriteBinary(data, opt)
}

func (inst *ImpAppRuntimeService) getLogFile() afs.Path {
	path := inst.AppDataService.GetAppDataDirectory()
	dir := inst.FileSystemService.Path(path)
	return dir.GetChild("runtime.json")
}

func (inst *ImpAppRuntimeService) backupSelf() error {

	enabled := inst.EnableBackupSelf
	if !enabled {
		return nil
	}

	info, err := inst.getAppRuntimeInfo()
	if err != nil {
		return err
	}

	src := inst.FileSystemService.Path(info.Path)
	wantSum := info.SHA256SUM
	wantSize := info.Size

	pathStr := inst.AppDataService.GetBackupExecutableFile(wantSum)
	dst := inst.FileSystemService.Path(pathStr)

	// check dst
	if dst.IsFile() {
		haveSize := dst.GetInfo().Length()
		if haveSize == wantSize {
			return nil // skip
		}
	}

	// check src
	if !src.IsFile() {
		return fmt.Errorf("src not a file")
	}
	if src.GetInfo().Length() != wantSize {
		return fmt.Errorf("bad src size")
	}

	// copy data
	haveSize, err := inst.copyFile(src, dst)
	if err != nil {
		return err
	}
	if haveSize != wantSize {
		return fmt.Errorf("bad file size")
	}

	return nil
}

func (inst *ImpAppRuntimeService) copyFile(src, dst afs.Path) (int64, error) {

	r, err := src.GetIO().OpenReader(nil)
	if err != nil {
		return 0, err
	}
	defer r.Close()

	opt := &afs.Options{Mkdirs: true, Create: true}

	w, err := dst.GetIO().OpenWriter(opt)
	if err != nil {
		return 0, err
	}
	defer w.Close()

	return io.Copy(w, r)
}

// ReadStartupLogs ...
func (inst *ImpAppRuntimeService) ReadStartupLogs() (*backup.StartupVO, error) {
	o := &backup.StartupVO{}
	err := inst.loadLogs(o)
	if err != nil {
		return nil, err
	}
	return o, nil
}

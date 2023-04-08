package backups

import (
	"context"
	"fmt"
	"strings"
	"time"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// ImpBackupService ...
type ImpBackupService struct {
	markup.Component `id:"DatabaseBackupService" class:"life"`

	AppDataService    service.AppDataService    `inject:"#AppDataService"`
	AppRuntimeService service.AppRuntimeService `inject:"#AppRuntimeService"`
	FilesysService    service.FileSystemService `inject:"#FileSystemService"`
	BackupDao         dao.Backup                `inject:"#wpm-database-backup-dao"`

	DoDump bool `inject:"${wpm.db.dump.enabled}"`
}

func (inst *ImpBackupService) _Impl() (service.DatabaseBackupService, application.LifeRegistry) {
	return inst, inst
}

// GetLifeRegistration ...
func (inst *ImpBackupService) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnStop: inst.onStop,
	}
}

func (inst *ImpBackupService) onStop() error {
	if inst.DoDump {
		return inst.dumpSelf()
	}
	return nil
}

func (inst *ImpBackupService) preparePathForExport(o *dto.Backup) (afs.Path, error) {

	n1 := o.BackupFileName
	n2 := o.BackupFilePath
	if n1 == "" && n2 == "" {
		return nil, fmt.Errorf("no file path")
	}
	name := n2
	if name == "" {
		name = n1
	}

	path := inst.getBackupDir()
	name = inst.addFileNameSuffixAuto(name, ".json")

	if strings.ContainsRune(name, '\\') || strings.ContainsRune(name, '/') {
		// as a path , nop
		path = path.GetFS().NewPath(name)
	} else {
		// as simple file name
		path = path.GetChild(name)
	}

	dir := path.GetParent()
	file := path

	if !dir.Exists() {
		dir.Mkdirs(nil)
	}

	if file.Exists() {
		return nil, fmt.Errorf("the file is exists, path=" + file.GetPath())
	}

	return file, nil
}

func (inst *ImpBackupService) addFileNameSuffixAuto(name, suffix string) string {
	n1 := strings.ToLower(name)
	s1 := strings.ToLower(suffix)
	if strings.HasSuffix(n1, s1) {
		return name
	}
	return name + suffix
}

// Export ...
func (inst *ImpBackupService) Export(c context.Context, o *dto.Backup) (*dto.Backup, error) {

	path, err := inst.preparePathForExport(o)
	if err != nil {
		return nil, err
	}

	err = inst.BackupDao.Export(path)
	if err != nil {
		return nil, err
	}

	info := path.GetInfo()
	t1 := info.CreatedAt()
	t2 := info.CreatedAt()

	o.BackupFileName = path.GetName()
	o.BackupFilePath = path.GetPath()
	o.CreatedAt = util.NewTime(t1)
	o.UpdatedAt = util.NewTime(t2)
	return o, nil
}

// Import ...
func (inst *ImpBackupService) Import(c context.Context, o *dto.Backup) (*dto.Backup, error) {
	path := o.BackupFilePath
	path = strings.TrimSpace(path)
	if path == "" {
		return nil, fmt.Errorf("no param:backup_file_path")
	}
	file := inst.FilesysService.Path(path)
	err := inst.BackupDao.Import(file)
	if err != nil {
		return nil, err
	}
	return o, nil
}

// ListAll ...
func (inst *ImpBackupService) ListAll(c context.Context) ([]*dto.Backup, error) {
	dir := inst.getBackupDir()
	src := inst.findAllJSONFileInDir(dir)
	dst := make([]*dto.Backup, 0)
	for _, item1 := range src {
		item2 := inst.makeBackupItem(item1)
		if item2 != nil {
			dst = append(dst, item2)
		}
	}
	return dst, nil
}

func (inst *ImpBackupService) makeBackupItem(file afs.Path) *dto.Backup {
	if file == nil {
		return nil
	}
	if !file.IsFile() {
		return nil
	}
	info := file.GetInfo()
	t1 := info.CreatedAt()
	t2 := info.UpdatedAt()
	item := &dto.Backup{
		BackupFileName: file.GetName(),
		BackupFilePath: file.GetPath(),
	}
	item.CreatedAt = util.NewTime(t1)
	item.UpdatedAt = util.NewTime(t2)
	return item
}

func (inst *ImpBackupService) findAllJSONFileInDir(dir afs.Path) []afs.Path {
	dst := make([]afs.Path, 0)
	if !dir.IsDirectory() {
		return dst
	}
	src := dir.ListChildren()
	for _, p := range src {
		name := strings.ToLower(p.GetName())
		if p.IsFile() && strings.HasSuffix(name, ".json") {
			dst = append(dst, p)
		}
	}
	return dst
}

func (inst *ImpBackupService) getBackupDir() afs.Path {
	path1 := inst.AppDataService.GetBackupDumpDirectory()
	path2 := inst.FilesysService.Path(path1)
	if !path2.Exists() {
		path2.Mkdirs(nil)
	}
	return path2
}

func (inst *ImpBackupService) dumpSelf() error {

	now := time.Now()
	name := "dump-" + now.Format(time.RFC3339)
	name = strings.ReplaceAll(name, ":", "")
	name = strings.ReplaceAll(name, "-", "_")
	name = strings.ReplaceAll(name, "+", "_")
	name = strings.ToLower(name)

	bak := &dto.Backup{
		BackupFileName: name,
		BackupFilePath: "",
	}

	file, err := inst.preparePathForExport(bak)
	if err != nil {
		return err
	}

	c := context.Background()
	_, err = inst.Export(c, bak)
	if err != nil {
		return err
	}

	vlog.Info("dump to file " + file.GetPath())
	return nil
}

// ExportDumpData ...
func (inst *ImpBackupService) ExportDumpData(c context.Context) error {
	h := &myDumpExportHandler{
		context:           c,
		AppRuntimeService: inst.AppRuntimeService,
		AppDataService:    inst.AppDataService,
		FileSystemService: inst.FilesysService,
	}
	return h.Export()
}

// ImportDumpData ...
func (inst *ImpBackupService) ImportDumpData(c context.Context, from, to util.Time) error {
	h := &myDumpImportHandler{
		context:               c,
		AppDataService:        inst.AppDataService,
		DatabaseBackupService: inst,
		FileSystemService:     inst.FilesysService,
	}
	return h.Import(from, to)
}

package backups

import (
	"context"
	"fmt"
	"strings"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

type myDumpExportHandler struct {
	context context.Context
}

func (inst *myDumpExportHandler) Export() error {
	return fmt.Errorf("no impl")
}

////////////////////////////////////////////////////////////////////////////////

type myDumpImportHandler struct {
	context               context.Context
	AppDataService        service.AppDataService
	FileSystemService     service.FileSystemService
	DatabaseBackupService service.DatabaseBackupService
}

func (inst *myDumpImportHandler) Import(from, to util.Time) error {
	ctx := inst.context
	file, err := inst.findLatestDumpFile(from, to)
	if err != nil {
		return err
	}
	o := &dto.Backup{}
	o.BackupFilePath = file.GetPath()
	_, err = inst.DatabaseBackupService.Import(ctx, o)
	return err
}

func (inst *myDumpImportHandler) findLatestDumpFile(from, to util.Time) (afs.Path, error) {
	dirPath := inst.AppDataService.GetBackupDumpDirectory()
	dir := inst.FileSystemService.Path(dirPath)
	if !dir.IsDirectory() {
		return nil, fmt.Errorf("no dump dir at  %v", dir.GetPath())
	}
	all := dir.ListChildren()
	for _, file := range all {
		if !file.IsFile() {
			continue
		}
		name := file.GetName()
		if !strings.HasSuffix(name, ".json") {
			continue
		}
		info := file.GetInfo()
		at := util.NewTime(info.CreatedAt())
		if from <= at && at <= to {
			return file, nil
		}
	}
	return nil, fmt.Errorf("no dump file right now")
}

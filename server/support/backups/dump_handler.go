package backups

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
)

type myDumpExportHandler struct {
	context context.Context

	AppRuntimeService service.AppRuntimeService
	AppDataService    service.AppDataService
	FileSystemService service.FileSystemService

	// var
	history []*entity.Executable
	self    *entity.Executable
	result  *entity.Executable
	tempExe afs.Path
}

func (inst *myDumpExportHandler) Export() error {

	steps := make([]func() error, 0)
	steps = append(steps, inst.loadThisRuntimeInfo)
	steps = append(steps, inst.loadRuntimeHistory)
	steps = append(steps, inst.scanRuntimeHistory)
	steps = append(steps, inst.prepareTempExe)
	steps = append(steps, inst.runTempExe)

	for _, fn := range steps {
		err := fn()
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *myDumpExportHandler) loadThisRuntimeInfo() error {

	sum, err := inst.AppRuntimeService.GetAppHash()
	if err != nil {
		return err
	}

	o := &entity.Executable{}
	o.CreatedAt = time.Now()
	o.SHA256SUM = sum

	inst.self = o
	return nil
}

func (inst *myDumpExportHandler) loadRuntimeHistory() error {
	logs, err := inst.AppRuntimeService.ReadStartupLogs()
	if err != nil {
		return nil
	}
	self := inst.self
	src := logs.Versions
	dst := make([]*entity.Executable, 0)
	for _, item := range src {
		if item == nil {
			continue
		} else if item.SHA256SUM == self.SHA256SUM {
			continue
		}
		dst = append(dst, item)
	}
	inst.history = dst
	inst.sortHistory()
	return nil
}

func (inst *myDumpExportHandler) scanRuntimeHistory() error {
	all := inst.history
	inst.result = nil
	for _, item := range all {
		sum := item.SHA256SUM
		// date := item.CreatedAt.Format(time.RFC3339)
		// vlog.Warn("history.item: ", sum, " : ", date)
		exeFilePath := inst.AppDataService.GetBackupExecutableFile(sum)
		file := inst.FileSystemService.Path(exeFilePath)
		if file.IsFile() {
			size := file.GetInfo().Length()
			if size == item.Size {
				item.Path = file.GetPath()
				inst.result = item
				break
			}
		}
	}
	if inst.result == nil {
		return fmt.Errorf("no available dump file")
	}
	return nil
}

func (inst *myDumpExportHandler) prepareTempExe() error {

	// src
	path1 := inst.result.Path
	size1 := inst.result.Size
	hash1 := inst.result.SHA256SUM
	file1 := inst.FileSystemService.Path(path1)

	// dst
	path2 := os.TempDir() + "/wpm/bin/" + hash1.String() + "/wpm-dump.exe"
	file2 := inst.FileSystemService.Path(path2)

	if file2.IsFile() {
		// check
		size2 := file2.GetInfo().Length()
		if size2 != size1 {
			return fmt.Errorf("bad wpm-dump executable size")
		}
		hash2, err := utils.ComputeFileSHA256sumAFS(file2)
		if err != nil {
			return err
		}
		if hash2 != hash1 {
			return fmt.Errorf("bad wpm-dump executable(1) sha-256")
		}
	} else {
		// copy
		opt := &afs.Options{Mkdirs: true, Create: true}
		size2, err := utils.CopyFile(file1, file2, opt)
		if err != nil {
			return err
		}
		if size2 != size1 {
			return fmt.Errorf("bad wpm-dump executable size")
		}
	}
	hash2, err := utils.ComputeFileSHA256sumAFS(file2)
	if err != nil {
		return err
	}
	if hash1 != hash2 {
		return fmt.Errorf("bad wpm-dump executable(2) sha-256")
	}
	inst.tempExe = file2
	return nil
}

func (inst *myDumpExportHandler) runTempExe() error {

	exe := inst.tempExe
	args := []string{
		"--application.profiles.active=dump",
	}

	cmd := exec.Command(exe.GetPath(), args...)
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return err
	}
	return cmd.Wait()
}

func (inst *myDumpExportHandler) sortHistory() {
	sort.Sort(inst)
}
func (inst *myDumpExportHandler) Less(i1, i2 int) bool {
	t1 := inst.history[i1].CreatedAt
	t2 := inst.history[i2].CreatedAt
	return t1.After(t2)
}
func (inst *myDumpExportHandler) Swap(i1, i2 int) {
	o1 := inst.history[i1]
	o2 := inst.history[i2]
	inst.history[i1] = o2
	inst.history[i2] = o1
}
func (inst *myDumpExportHandler) Len() int {
	return len(inst.history)
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
	} else if err == nil && file == nil {
		return nil // no dump file right now
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
	// return nil, nil
}

package mediae

import (
	"context"
	"fmt"
	"io/fs"
	"strconv"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
)

type myMediaImporter struct {

	// context
	context    context.Context
	parent     *MediaServiceImpl
	fs         service.FileSystemService
	httpclient service.HTTPClientService

	// var
	m1        *dto.Media
	m2        *dto.Media
	fileDest  afs.Path
	fileTemp  afs.Path
	skipFetch bool
	wantSize  int64
	wantSum   util.Hex
}

func (inst *myMediaImporter) Import(m1 *dto.Media, opt *service.MediaOptions) (*dto.Media, error) {

	m2 := &dto.Media{}
	inst.m1 = m1
	inst.m2 = m2
	steps := make([]func() error, 0)

	steps = append(steps, inst.prepare)
	steps = append(steps, inst.computeDestFilePath)
	steps = append(steps, inst.checkDestIsReady)
	steps = append(steps, inst.computeTempFilePath)

	steps = append(steps, inst.fetchFromSource)
	steps = append(steps, inst.checkTempFile)
	steps = append(steps, inst.moveTempToDest)

	steps = append(steps, inst.fillMedia2)

	defer func() {
		err := inst.clean()
		if err != nil {
			vlog.Warn(err)
		}
	}()

	for _, fn := range steps {
		err := fn()
		if err != nil {
			return nil, err
		}
	}

	return m2, nil
}

func (inst *myMediaImporter) prepare() error {
	m1 := inst.m1
	m2 := inst.m2
	if m1 == nil || m2 == nil {
		return fmt.Errorf("no media object")
	}
	inst.wantSize = m1.FileSize
	inst.wantSum = m1.SHA256SUM
	return nil
}

func (inst *myMediaImporter) computeDestFilePath() error {
	ctx := inst.context
	mdir, err := inst.parent.getSystemMediaDir(ctx)
	if err != nil {
		return err
	}
	m := inst.m1
	file, err := makeMediaLocalPath(m, mdir)
	if err != nil {
		return err
	}
	inst.fileDest = file
	return nil
}

func (inst *myMediaImporter) computeTempFilePath() error {
	dest := inst.fileDest
	if dest == nil {
		return fmt.Errorf("no dest file path")
	}
	name := dest.GetName()
	dir := dest.GetParent()
	n := util.Now().Int64()
	t := strconv.FormatInt(n, 10)
	tmp := dir.GetChild(name + "-" + t + ".tmp~")
	inst.fileTemp = tmp
	return nil
}

func (inst *myMediaImporter) isDestReady() (bool, error) {

	file := inst.fileDest
	if file == nil {
		return false, fmt.Errorf("dest media file is nil")
	}

	if !file.Exists() {
		return false, fmt.Errorf("no dest media file")
	}

	haveSize := file.GetInfo().Length()
	if haveSize != inst.wantSize {
		return false, fmt.Errorf("bad size")
	}

	haveHash, err := utils.ComputeFileSHA256sumAFS(file)
	if err != nil {
		return false, err
	}
	if haveHash != inst.wantSum {
		return false, fmt.Errorf("bad sha-256 sum")
	}

	return true, nil
}

func (inst *myMediaImporter) checkDestIsReady() error {
	ready, _ := inst.isDestReady()
	if ready {
		inst.skipFetch = true // 如果目标已就绪，跳过 fetch 步骤
	}
	return nil
}

func (inst *myMediaImporter) checkTempFile() error {

	if inst.skipFetch {
		return nil
	}

	file := inst.fileTemp
	haveSize := file.GetInfo().Length()
	haveHash, err := utils.ComputeFileSHA256sumAFS(file)
	if err != nil {
		return err
	}

	if haveSize != inst.wantSize {
		return fmt.Errorf("bad media size")
	}

	if haveHash != inst.wantSum {
		return fmt.Errorf("bad media sha-256 sum")
	}

	return nil
}

func (inst *myMediaImporter) moveTempToDest() error {

	if inst.skipFetch {
		return nil
	}

	opt := &afs.Options{Create: true}
	dst := inst.fileDest
	src := inst.fileTemp
	return src.CopyTo(dst, opt)
}

func (inst *myMediaImporter) clean() error {
	file := inst.fileTemp
	if file == nil {
		return nil
	}
	if !file.IsFile() {
		return nil
	}
	return file.Delete()
}

func (inst *myMediaImporter) fetchFromSource() error {

	if inst.skipFetch {
		return nil
	}

	ctx := inst.context
	url := inst.m1.Source

	if utils.IsAbsolutePath(url) {
		url = "file:///" + url
	}

	data, _, err := inst.httpclient.FetchBinary(ctx, url, nil)
	if err != nil {
		return err
	}
	tmp := inst.fileTemp
	opt := &afs.Options{
		Mkdirs: true, Create: true, Permission: fs.ModePerm,
	}
	return tmp.GetIO().WriteBinary(data, opt)
}

func (inst *myMediaImporter) fillMedia2() error {

	dst := inst.fileDest
	o1 := inst.m1
	o2 := inst.m2

	url, err := makeMediaWebPath(o1)
	if err != nil {
		return err
	}

	if o1.Bucket != "" {
		o2.Bucket = o1.Bucket
	} else {
		o2.Bucket = "default"
	}

	// o2.Source = dst.GetPath()
	o2.Source = o1.Source
	o2.URL = url
	o2.FileSize = dst.GetInfo().Length()
	o2.ContentType = o1.ContentType
	o2.Name = o1.Name
	o2.Label = o1.Label
	o2.SHA256SUM = inst.wantSum

	return nil
}

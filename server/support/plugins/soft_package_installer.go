package plugins

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

type myPakcageInstaller struct {
	context      context.Context
	HTTPClient   service.HTTPClientService
	HTTPClientEx service.HTTPClientExService

	ContentTypeSer    service.ContentTypeService
	MediaSer          service.MediaService
	IntentTemplateSer service.IntentTemplateService
	ExecutableSer     service.ExecutableService

	installation dxo.InstallationID
	pack         *dto.SoftwarePackage
	res          *vo.Online
}

func (inst *myPakcageInstaller) Install(p *dto.SoftwarePackage) (*dto.SoftwarePackage, error) {

	inst.pack = p

	steps := make([]func() error, 0)
	steps = append(steps, inst.fetchJSON)
	steps = append(steps, inst.fetchZip)
	steps = append(steps, inst.applyRes)

	for _, step := range steps {
		err := step()
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (inst *myPakcageInstaller) fetchJSON() error {
	ctx := inst.context
	url := inst.pack.ResourceURL
	if url == "" {
		inst.res = &vo.Online{}
		return nil
	}
	doc, err := inst.HTTPClientEx.FetchOnlineDoc(ctx, url, nil)
	if err != nil {
		return err
	}
	inst.res = doc
	return nil
}

func (inst *myPakcageInstaller) fetchZip() error {

	ctx := inst.context
	url := inst.pack.DownloadURL
	wantSize := inst.pack.Size
	wantHash := inst.pack.SHA256SUM

	// 检查是否包含有效的压缩包
	if url == "" || wantSize < 1 || wantHash == "" {
		return nil
	}

	data, resp, err := inst.HTTPClient.FetchBinary(ctx, url, nil)
	if err != nil {
		return err
	}

	// check size
	haveSize1 := resp.ContentLength
	haveSize2 := int64(len(data))
	if haveSize1 != wantSize || haveSize2 != wantSize {
		return fmt.Errorf("bad package size")
	}

	// check sum
	wantSum := wantHash.Bytes()
	haveSum := sha256.Sum256(data)
	if bytes.Compare(wantSum, haveSum[:]) != 0 {
		return fmt.Errorf("bad package sha-256 sum")
	}

	return nil
}

func (inst *myPakcageInstaller) applyRes() error {

	errlist := &utils.ErrorList{}
	r := inst.res

	inst.applyExecutables(r.Executables, errlist)
	inst.applyIntentTemplates(r.IntentTemplates, errlist)
	inst.applyMediae(r.Mediae, errlist)
	inst.applyContentTypes(r.ContentTypes, errlist)

	allerr := errlist.All()
	for _, err := range allerr {
		if err == nil {
			continue
		}
		vlog.Warn(err.Error())
	}

	return nil
}

func (inst *myPakcageInstaller) applyIntentTemplates(list []*dto.IntentTemplate, elist *utils.ErrorList) {
	ctx := inst.context
	ser := inst.IntentTemplateSer
	iid := inst.installation
	for _, item := range list {
		item.Installation = iid
		_, err := ser.Insert(ctx, item)
		elist.Append(err)
	}
}

func (inst *myPakcageInstaller) applyExecutables(list []*dto.Executable, elist *utils.ErrorList) {
	ctx := inst.context
	ser := inst.ExecutableSer
	iid := inst.installation
	for _, item := range list {
		item.Installation = iid
		_, err := ser.Insert(ctx, item, nil)
		elist.Append(err)
	}
}

func (inst *myPakcageInstaller) applyMediae(list []*dto.Media, elist *utils.ErrorList) {
	ctx := inst.context
	ser := inst.MediaSer
	iid := inst.installation
	for _, item := range list {
		item.Installation = iid
		_, err := ser.Insert(ctx, item, nil)
		elist.Append(err)
	}
}

func (inst *myPakcageInstaller) applyContentTypes(list []*dto.ContentType, elist *utils.ErrorList) {
	ctx := inst.context
	ser := inst.ContentTypeSer
	iid := inst.installation
	for _, item := range list {
		item.Installation = iid
		_, err := ser.Insert(ctx, item)
		elist.Append(err)
	}
}

package plugins

import (
	"bytes"
	"context"
	"fmt"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/components/packs"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

type myPakcageInstaller struct {
	context context.Context

	HTTPClient         service.HTTPClientService
	HTTPClientEx       service.HTTPClientExService
	ContentTypeSer     service.ContentTypeService
	MediaSer           service.MediaService
	IntentTemplateSer  service.IntentTemplateService
	ExecutableSer      service.ExecutableService
	SoftwarePackageSer service.SoftwarePackageService
	AppDataSer         service.AppDataService

	installation dxo.InstallationID
	pack         *dto.SoftwarePackage
	res          *vo.Online
	pic          *packs.InstallingContext
}

func (inst *myPakcageInstaller) Install(p *dto.SoftwarePackage) (*dto.SoftwarePackage, error) {

	inst.pack = p

	steps := make([]func() error, 0)
	steps = append(steps, inst.prepare)
	steps = append(steps, inst.fetch)
	steps = append(steps, inst.apply)

	for _, step := range steps {
		err := step()
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (inst *myPakcageInstaller) prepare() error {

	iid := inst.installation
	pack := inst.pack
	sum := pack.SHA256SUM
	ads := inst.AppDataSer
	id := sum.String()

	pack.Installation = iid

	pic := &packs.InstallingContext{}
	pic.Context = inst.context
	pic.Action = "install"
	pic.Pack = pack
	pic.PackType = pack.ContentType
	pic.Installation = iid

	pic.PackDir = ads.GetPath(&service.GetPathOptions{Type: "plug-ins", ID: id})
	pic.PackFile = ads.GetPath(&service.GetPathOptions{Type: "packages-cache", ID: id})
	pic.Root = ads.GetRoot()

	inst.pic = pic
	return nil
}

// func (inst *myPakcageInstaller) fetchJSON() error {
// 	ctx := inst.context
// 	url := inst.pack.ResourceURL
// 	if url == "" {
// 		inst.res = &vo.Online{}
// 		return nil
// 	}
// 	doc, err := inst.HTTPClientEx.FetchOnlineDoc(ctx, url, nil)
// 	if err != nil {
// 		return err
// 	}
// 	inst.res = doc
// 	return nil
// }

func (inst *myPakcageInstaller) fetch() error {

	ctx := inst.context
	url := inst.pack.DownloadURL
	wantSize := inst.pack.Size
	wantHash := inst.pack.SHA256SUM

	// 检查是否包含有效的压缩包
	if url == "" || wantSize < 1 || wantHash == "" {
		return nil
	}

	haveSize1 := int64(0)
	file2 := inst.pic.PackFile

	if !file2.Exists() {
		vlog.Info("fetch package from ", url)
		opt := &service.HTTPClientOptions{}
		opt.FileOptions = &afs.Options{
			Mkdirs: true,
			Create: true,
			Write:  true,
			File:   true,
		}
		resp, err := inst.HTTPClient.FetchToFile(ctx, url, file2, opt)
		if err != nil {
			return err
		}
		haveSize1 = resp.ContentLength
	} else if file2.IsFile() {
		haveSize1 = file2.GetInfo().Length()
	}

	// check size
	haveSize2 := file2.GetInfo().Length()
	if haveSize1 != wantSize || haveSize2 != wantSize {
		return fmt.Errorf("bad package size")
	}

	// check sum
	wantSum := wantHash.Bytes()
	haveSum, _ := utils.ComputeFileSHA256sumAFS(file2)
	if bytes.Compare(wantSum, haveSum.Bytes()) != 0 {
		return fmt.Errorf("bad package sha-256 sum")
	}

	return nil
}

// func (inst *myPakcageInstaller) applyRes() error {

// 	errlist := &utils.ErrorList{}
// 	r := inst.res

// 	inst.applyExecutables(r.Executables, errlist)
// 	inst.applyIntentTemplates(r.IntentTemplates, errlist)
// 	inst.applyMediae(r.Mediae, errlist)
// 	inst.applyContentTypes(r.ContentTypes, errlist)

// 	allerr := errlist.All()
// 	for _, err := range allerr {
// 		if err == nil {
// 			continue
// 		}
// 		vlog.Warn(err.Error())
// 	}

// 	return nil
// }

// func (inst *myPakcageInstaller) applyIntentTemplates(list []*dto.IntentTemplate, elist *utils.ErrorList) {
// 	ctx := inst.context
// 	ser := inst.IntentTemplateSer
// 	iid := inst.installation
// 	for _, item := range list {
// 		item.Installation = iid
// 		_, err := ser.Insert(ctx, item)
// 		elist.Append(err)
// 	}
// }

// func (inst *myPakcageInstaller) applyExecutables(list []*dto.Executable, elist *utils.ErrorList) {
// 	ctx := inst.context
// 	ser := inst.ExecutableSer
// 	iid := inst.installation
// 	for _, item := range list {
// 		item.Installation = iid
// 		_, err := ser.Insert(ctx, item, nil)
// 		elist.Append(err)
// 	}
// }

// func (inst *myPakcageInstaller) applyMediae(list []*dto.Media, elist *utils.ErrorList) {
// 	ctx := inst.context
// 	ser := inst.MediaSer
// 	iid := inst.installation
// 	for _, item := range list {
// 		item.Installation = iid
// 		_, err := ser.Insert(ctx, item, nil)
// 		elist.Append(err)
// 	}
// }

// func (inst *myPakcageInstaller) applyContentTypes(list []*dto.ContentType, elist *utils.ErrorList) {
// 	ctx := inst.context
// 	ser := inst.ContentTypeSer
// 	iid := inst.installation
// 	for _, item := range list {
// 		item.Installation = iid
// 		_, err := ser.Insert(ctx, item)
// 		elist.Append(err)
// 	}
// }

func (inst *myPakcageInstaller) apply() error {
	pic := inst.pic
	pm := inst.SoftwarePackageSer.GetPacksManger()
	installer, err := pm.GetInstaller(pic)
	if err != nil {
		return err
	}
	return installer.Install(pic)
}

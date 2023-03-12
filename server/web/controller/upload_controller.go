package controller

import (
	"net/http"
	"os"
	"strconv"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/gin-gonic/gin"
)

// UploadController Upload 控制器
type UploadController struct {
	markup.RestController `class:"rest-controller"`

	// PlatformService service.PlatformService `inject:"#PlatformService"`
	// ProfileService  service.ProfileService  `inject:"#ProfileService"`

	FileSystemService service.FileSystemService `inject:"#FileSystemService"`
	Responder         glass.MainResponder       `inject:"#glass-main-responder"`
}

func (inst *UploadController) _Impl() glass.Controller {
	return inst
}

// Init 初始化
func (inst *UploadController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("/upload")
	ec.Handle(http.MethodPost, "", inst.handleUpload)
	return nil
}

func (inst *UploadController) handleUpload(c *gin.Context) {
	req := &myUploadRequest{
		parent: inst,
		gc:     c,
	}
	err := req.open()
	if err == nil {
		err = req.upload()
	}
	req.send(err)
}

////////////////////////////////////////////////////////////////////////////////

type myUploadRequest struct {
	parent *UploadController
	gc     *gin.Context
	save   bool
	body2  vo.Media
}

func (inst *myUploadRequest) send(err error) {
	ctx := inst.gc
	status := inst.body2.Status
	body := &inst.body2
	resp := &glass.Response{
		Context: ctx,
		Data:    body,
		Error:   err,
		Status:  status,
	}
	inst.parent.Responder.Send(resp)
}

func (inst *myUploadRequest) open() error {
	return nil
}

func (inst *myUploadRequest) makeTempFile() afs.Path {

	now := util.Now().Int64()
	ts := strconv.FormatInt(now, 10)
	path := os.TempDir() + "/bitwormhole/wpm/media/" + ts

	file := inst.parent.FileSystemService.FS().NewPath(path)
	dir := file.GetParent()
	if !dir.Exists() {
		dir.Mkdirs(nil)
	}

	return file
}

func (inst *myUploadRequest) cleanTempFile(file afs.Path) {
	if file == nil {
		return
	}
	if !file.IsFile() {
		return
	}
	if inst.save {
		return
	}
	file.Delete()
}

func (inst *myUploadRequest) uploadMedia(tmp afs.Path) (*dto.Media, error) {

	c := inst.gc
	// name := c.PostForm("name")
	dst := tmp.GetPath()

	file, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}

	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		return nil, err
	}

	sum, err := utils.ComputeFileSHA256sumAFS(tmp)
	if err != nil {
		return nil, err
	}

	ctype := file.Header.Get("Content-Type")

	media := &dto.Media{}
	media.URL = ""
	media.LocalFilePath = tmp.GetPath()
	media.Bucket = "temp"
	media.ContentType = ctype
	media.FileSize = tmp.GetInfo().Length()
	media.SHA256SUM = sum
	media.Name = file.Filename
	media.Label = file.Filename

	return media, nil
}

func (inst *myUploadRequest) upload() error {
	tmp := inst.makeTempFile()
	defer func() {
		inst.cleanTempFile(tmp)
	}()
	media, err := inst.uploadMedia(tmp)
	if err != nil {
		return err
	}
	inst.body2.Mediae = []*dto.Media{media}
	inst.save = true
	return nil
}

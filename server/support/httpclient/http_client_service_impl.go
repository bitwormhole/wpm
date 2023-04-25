package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
)

// ImpHTTPClientService ....
type ImpHTTPClientService struct {
	markup.Component `id:"HTTPClientService"`

	AC application.Context       `inject:"context"`
	FS service.FileSystemService `inject:"#FileSystemService"`

	MaxContentLength int `inject:"${wpm.httpclient.max-content-length}"`
}

func (inst *ImpHTTPClientService) _Impl() service.HTTPClientService {
	return inst
}

// FetchText 通过 HTTP 协议获取文本
func (inst *ImpHTTPClientService) FetchText(ctx context.Context, url string, opt *service.HTTPClientOptions) (string, *service.HTTPClientResult, error) {
	data, res, err := inst.fetch(ctx, url, opt)
	if err != nil {
		return "", res, err
	}
	text := string(data)
	return text, res, err
}

// FetchJSON 通过 HTTP 协议获取 JSON 对象
func (inst *ImpHTTPClientService) FetchJSON(ctx context.Context, url string, obj any, opt *service.HTTPClientOptions) (*service.HTTPClientResult, error) {

	data, res, err := inst.fetch(ctx, url, opt)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, obj)
	if err != nil {
		return res, err
	}

	return res, nil
}

// FetchBinary 通过 HTTP 协议获取二进制数据
func (inst *ImpHTTPClientService) FetchBinary(ctx context.Context, url string, opt *service.HTTPClientOptions) ([]byte, *service.HTTPClientResult, error) {
	return inst.fetch(ctx, url, opt)
}

func (inst *ImpHTTPClientService) prepareOptions(opt *service.HTTPClientOptions) *service.HTTPClientOptions {
	if opt == nil {
		opt = &service.HTTPClientOptions{}
	}
	if opt.MaxContentLength < 100 {
		opt.MaxContentLength = int64(inst.MaxContentLength)
	}
	return opt
}

func (inst *ImpHTTPClientService) isWebURL(url string) bool {
	return strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://")
}

func (inst *ImpHTTPClientService) isFileURL(url string) bool {
	return strings.HasPrefix(url, "file:/")
}

func (inst *ImpHTTPClientService) isResURL(url string) bool {
	accepts := []string{"res://", "resource://", "resources://"}
	for _, prefix := range accepts {
		if strings.HasPrefix(url, prefix) {
			return true
		}
	}
	return false
}

func (inst *ImpHTTPClientService) fetchFromRes(ctx context.Context, url string, opt *service.HTTPClientOptions) ([]byte, *service.HTTPClientResult, error) {
	res := &service.HTTPClientResult{}
	data, err := inst.AC.GetResources().GetBinary(url)
	if err == nil {
		size := len(data)
		res.Status = http.StatusOK
		res.StatusText = "OK"
		res.ContentLength = int64(size)
		res.ContentType = "application/octet-stream"
	} else {
		res.StatusText = err.Error()
	}
	return data, res, err
}

func (inst *ImpHTTPClientService) fetchFromLocalFS(ctx context.Context, url string, opt *service.HTTPClientOptions) ([]byte, *service.HTTPClientResult, error) {

	const prefix = "file:/"
	path := url
	if strings.HasPrefix(url, prefix) {
		path = url[len(prefix):]
	}

	errNoFile := fmt.Errorf("no file at [%v]", url)
	file := inst.FS.Path(path)
	if file == nil {
		return nil, nil, errNoFile
	}
	if !file.IsFile() {
		return nil, nil, errNoFile
	}

	res := &service.HTTPClientResult{}
	data, err := file.GetIO().ReadBinary(nil)
	if err == nil {
		size := len(data)
		res.Status = http.StatusOK
		res.StatusText = "OK"
		res.ContentLength = int64(size)
		res.ContentType = "application/octet-stream"
	} else {
		res.StatusText = err.Error()
	}
	return data, res, err
}

func (inst *ImpHTTPClientService) fetchFromWeb(ctx context.Context, url string, opt *service.HTTPClientOptions) ([]byte, *service.HTTPClientResult, error) {

	res := &service.HTTPClientResult{}

	resp, err := http.Get(url)
	if err != nil {
		return nil, res, err
	}

	body := resp.Body
	defer func() {
		if body != nil {
			body.Close()
		}
	}()

	res.Status = resp.StatusCode
	res.StatusText = resp.Status
	res.ContentLength = resp.ContentLength
	res.ContentType = resp.Header.Get("Content-Type")

	if res.Status != http.StatusOK {
		return nil, res, fmt.Errorf("HTTP " + res.StatusText)
	}

	if res.ContentLength > opt.MaxContentLength {
		return nil, res, fmt.Errorf("the HTTP content-length is too large, limit is:%v", opt.MaxContentLength)
	}

	data, err := io.ReadAll(body)
	if err != nil {
		return nil, res, err
	}

	return data, res, nil
}

func (inst *ImpHTTPClientService) fetch(ctx context.Context, url string, opt *service.HTTPClientOptions) ([]byte, *service.HTTPClientResult, error) {

	opt = inst.prepareOptions(opt)

	if inst.isResURL(url) {
		return inst.fetchFromRes(ctx, url, opt)
	} else if inst.isFileURL(url) {
		return inst.fetchFromLocalFS(ctx, url, opt)
	} else if inst.isWebURL(url) {
		return inst.fetchFromWeb(ctx, url, opt)
	}

	return nil, nil, fmt.Errorf("bad URL [%v]", url)
}

// FetchToStream ...
func (inst *ImpHTTPClientService) FetchToStream(ctx context.Context, url string, dst io.Writer, opt *service.HTTPClientOptions) (*service.HTTPClientResult, error) {
	// fetch binary
	data, res, err := inst.fetch(ctx, url, opt)
	if err != nil {
		return res, err
	}
	_, err = dst.Write(data)
	return res, err
}

// FetchToFile ...
func (inst *ImpHTTPClientService) FetchToFile(ctx context.Context, url string, dst afs.Path, opt *service.HTTPClientOptions) (*service.HTTPClientResult, error) {

	opt = inst.prepareOptions(opt)
	op2 := opt.FileOptions

	// fetch binary
	data, res, err := inst.fetch(ctx, url, opt)
	if err != nil {
		return res, err
	}
	err = dst.GetIO().WriteBinary(data, op2)
	return res, err
}

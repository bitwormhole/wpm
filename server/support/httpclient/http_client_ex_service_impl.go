package httpclient

import (
	"context"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// ImpHTTPClientEx ...
type ImpHTTPClientEx struct {
	markup.Component `id:"HTTPClientExService"`

	HTTPClientService service.HTTPClientService `inject:"#HTTPClientService"`
}

func (inst *ImpHTTPClientEx) _Impl() service.HTTPClientExService {
	return inst
}

// FetchOnlineDoc ...
func (inst *ImpHTTPClientEx) FetchOnlineDoc(ctx context.Context, url string, opt *service.HTTPClientOptions) (*vo.Online, error) {
	doc := &vo.Online{}
	_, err := inst.HTTPClientService.FetchJSON(ctx, url, doc, opt)
	if err != nil {
		return nil, err
	}

	head := &doc.Head
	head.URL = url

	inst.fillExecutables(head, doc.Executables)
	inst.fillMediae(head, doc.Mediae)
	inst.fillPackages(head, doc.Packages)

	return doc, nil
}

func (inst *ImpHTTPClientEx) fillExecutables(head *vo.BaseHead, list []*dto.Executable) {
	for _, item := range list {
		item.Referer = head.URL
		item.Namespace = inst.normalizeNamespace(head, item.Namespace)
		// item.IconURL = inst.resolveURL(head, item.IconURL)
	}
}

func (inst *ImpHTTPClientEx) fillMediae(head *vo.BaseHead, list []*dto.Media) {
	for _, item := range list {
		item.Referer = head.URL
	}
}

func (inst *ImpHTTPClientEx) fillPackages(head *vo.BaseHead, list []*dto.SoftwarePackage) {
	for _, item := range list {
		item.Referer = head.URL
		item.Namespace = inst.normalizeNamespace(head, item.Namespace)
		item.DownloadURL = inst.resolveURL(head, item.DownloadURL)
		item.ResourceURL = inst.resolveURL(head, item.ResourceURL)
		item.WebPageURL = inst.resolveURL(head, item.WebPageURL)
		item.Icon = inst.resolveURL(head, item.Icon)
	}
}

func (inst *ImpHTTPClientEx) normalizeNamespace(head *vo.BaseHead, ns string) string {
	if ns != "" {
		return ns
	}
	ns = head.Namespace
	if ns != "" {
		return ns
	}
	return "default"
}

func (inst *ImpHTTPClientEx) resolveURL(head *vo.BaseHead, url string) string {
	if url == "" {
		return ""
	}
	if head == nil {
		return url
	}
	r := &utils.URLResolver{}
	return r.ForDocument(head.URL).Resolve(url)
}

package service

import (
	"context"
)

// HTTPClientService ...
type HTTPClientService interface {

	// 通过 HTTP 协议获取文本
	FetchText(ctx context.Context, url string, opt *HTTPClientOptions) (string, *HTTPClientResult, error)

	// 通过 HTTP 协议获取二进制数据
	FetchBinary(ctx context.Context, url string, opt *HTTPClientOptions) ([]byte, *HTTPClientResult, error)

	// 通过 HTTP 协议获取 JSON 对象
	FetchJSON(ctx context.Context, url string, obj any, opt *HTTPClientOptions) (*HTTPClientResult, error)
}

// HTTPClientOptions ...
type HTTPClientOptions struct {
	MaxContentLength int64
}

// HTTPClientResult ...
type HTTPClientResult struct {
	ContentLength int64
	ContentType   string
	Status        int
	StatusText    string
}

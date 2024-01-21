package gui

import (
	"context"

	"github.com/bitwormhole/wpm/common/classes/about"
	"github.com/bitwormhole/wpm/server/web/vo"
	"github.com/starter-go/application"
	"github.com/starter-go/browsers"
)

// BrowserLauncher ...
type BrowserLauncher struct {

	//starter:component

	WebService   browsers.Service //starter:inject("#")
	AboutService about.Service    //starter:inject("#")

}

func (inst *BrowserLauncher) _impl() application.Lifecycle {
	return inst
}

// Life ...
func (inst *BrowserLauncher) Life() *application.Life {
	return &application.Life{
		OnStart: inst.start,
	}
}

func (inst *BrowserLauncher) start() error {
	ctx := context.Background()
	url := inst.getURL(ctx)
	i := &browsers.Intent{
		Context: ctx,
		URL:     url,
	}
	return inst.WebService.Run(i)
}

func (inst *BrowserLauncher) getURL(ctx context.Context) string {
	dst := new(vo.About)
	err := inst.AboutService.GetAboutInfo(ctx, dst)
	url := "https://bitwormhole.com/error/bad-wpm-web-url"
	if err == nil {
		url = dst.WebURL
	}
	return url
}

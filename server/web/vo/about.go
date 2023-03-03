package vo

import "github.com/bitwormhole/wpm/server/web/dto"

// About ...
type About struct {
	Base

	Name      string `json:"name"`
	Title     string `json:"title"`
	Copyright string `json:"copyright"`
	Profile   string `json:"profile"`
	User      string `json:"user"`
	Home      string `json:"home"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
	WebURL    string `json:"weburl"`

	// the main module
	Module     string     `json:"module"`
	Version    string     `json:"version"`
	Revision   int        `json:"revision"`
	MainModule dto.Module `json:"main_module"`

	Modules []*dto.Module `json:"modules"`
}

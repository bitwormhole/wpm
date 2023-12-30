package dto

import "github.com/starter-go/base/lang"

// Settings 表示一个命令模板
type Settings struct {
	// ID dxo.ExampleID `json:"id"`
	// Base

	// example:
	Value1 string `json:"value1"`
	Value2 int    `json:"value2"`

	SetupDone bool `json:"setup_done"`

	IgnorePackageVersion  string   `json:"ignore.package.version"`
	IgnorePackageRevision int      `json:"ignore.package.revision"`
	IgnorePackageSum      lang.Hex `json:"ignore.package.sha256sum"`
}

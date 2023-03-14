package dto

import "github.com/bitwormhole/starter/util"

// Settings 表示一个命令模板
type Settings struct {
	// ID dxo.ExampleID `json:"id"`
	// Base

	// example:
	Value1 string `json:"value1"`
	Value2 int    `json:"value2"`
	Value3 bool   `json:"value3"`

	IgnorePackageVersion  string   `json:"ignore.package.version"`
	IgnorePackageRevision int      `json:"ignore.package.revision"`
	IgnorePackageSum      util.Hex `json:"ignore.package.sha256sum"`
}

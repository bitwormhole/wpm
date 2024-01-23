package dto

import "github.com/bitwormhole/wpm/common/objects/dxo"

// ContentType 表示一个命令模板
type ContentType struct {
	ID  dxo.ContentTypeID  `json:"id"`
	URN dxo.ContentTypeURN `json:"urn"`

	Base

	TypeName    dxo.ContentTypeName `json:"type"`
	Patterns    dxo.StringList      `json:"patterns"`
	Label       string              `json:"label"`
	Description string              `json:"description"`
	Icon        string              `json:"icon"`

	Priority  int  `json:"priority"` // 优先级，数值越高越先处理
	AsFile    bool `json:"as_file"`
	AsDir     bool `json:"as_dir"`
	AsProject bool `json:"as_project"`
}

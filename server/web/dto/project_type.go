package dto

import "github.com/bitwormhole/wpm/server/data/dxo"

// ProjectType 表示一个命令模板
type ProjectType struct {
	ID  dxo.ProjectTypeID  `json:"id"`
	URN dxo.ProjectTypeURN `json:"urn"`

	Base

	TypeName    dxo.ProjectTypeName `json:"type"`
	Pattern     string              `json:"pattern"`
	Label       string              `json:"label"`
	Description string              `json:"description"`
	Icon        string              `json:"icon"`

	Priority  int  `json:"priority"` // 优先级，数值越高越先处理
	AsFile    bool `json:"as_file"`
	AsDir     bool `json:"as_dir"`
	AsProject bool `json:"as_project"`
}

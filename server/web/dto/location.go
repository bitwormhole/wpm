package dto

import "github.com/bitwormhole/wpm/server/data/dxo"

// Location 表示一个命令模板
type Location struct {
	ID dxo.LocationID `json:"id"`
	Base

	Path string `json:"path"` // the normalized full-path

	Class  dxo.LocationClass `json:"class"`
	AsFile bool              `json:"as_file"` // 表示该路径可以是一个普通文件
	AsDir  bool              `json:"as_dir"`  // 表示该路径可以是一个目录

	// [已废弃] 以下字段仅当 Class==LocationProject 时有用
	// TypeID   dxo.ProjectTypeID   `json:"type_id"`
	// TypeKey  dxo.ProjectTypeKey  `json:"type_key"`
	// TypeName dxo.ProjectTypeName `json:"type_name"`
}

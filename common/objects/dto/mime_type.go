package dto

import "github.com/bitwormhole/wpm/common/objects/dxo"

// MIMEType 表示一个 MIME 类型
type MIMEType struct {
	ID dxo.MIMETypeID `json:"id"`
	Base

	TypeName       string `json:"type"`
	FileNameSuffix string `json:"suffix"`
}

package entity

import (
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/starter-go/security-gorm/rbacdb"
	"gorm.io/gorm"
)

// Base ...
type Base struct {
	// UUID dxo.UUID `gorm:"index:,unique"`

	// ID        uint           `gorm:"primaryKey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`

	// Committer dxo.UserID
	// Creator   dxo.UserID
	// Owner     dxo.UserID

	rbacdb.BaseEntity

	Referer string // a URL, refer to owner document of this entity

	Installation dxo.InstallationID // 用来跟踪软件包安装资源项 ,  >0 表示已安装
}

// PrepareInsert ...
func (inst *Base) PrepareInsert() {
	if inst.Installation == 0 {
		inst.Installation = -1
	}
}

// HasDeletedAt 判断是否被软删除
func HasDeletedAt(at gorm.DeletedAt) bool {
	const limit = 3600 * 24 * 7
	n := at.Time.Unix()
	return n > limit
}

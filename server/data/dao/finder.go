package dao

import (
	"fmt"

	"github.com/starter-go/rbac"
	"gorm.io/gorm"
)

// Finder 是一个通过 gorm 查询 DB 的便捷工具
type Finder struct {
	DB          *gorm.DB
	ItemModel   any
	ListModel   any
	WantAll     bool
	Pagination  *rbac.Pagination // 分页参数
	QueryString string           // 查询条件（表达式）
	QueryParams []any            // 查询条件（参数）
}

// Find ...
func (inst Finder) Find() error {

	db := inst.DB
	item := inst.ItemModel
	list := inst.ListModel
	page := inst.Pagination
	query := inst.QueryString
	params := inst.QueryParams
	all := inst.WantAll

	// check params
	if db == nil {
		return fmt.Errorf("finder.param: db is nil")
	}
	if item == nil {
		return fmt.Errorf("finder.param: item is nil")
	}
	if list == nil {
		return fmt.Errorf("finder.param: list is nil")
	}
	if all {
		return inst.findAll()
	}
	db = db.Model(item)

	// query conditions
	if len(query) > 0 && params != nil {
		db = db.Where(query, params...)
	}

	// page & count
	if page != nil {
		// count
		count := int64(0)
		db.Count(&count)
		page.Total = count
		// pages
		limit := int(page.Limit())
		offset := int(page.Offset())
		db = db.Limit(limit).Offset(offset)
	}

	// find
	res := db.Find(list)
	return res.Error
}

func (inst Finder) findAll() error {

	db := inst.DB
	item := inst.ItemModel
	list := inst.ListModel
	page := inst.Pagination

	// find
	res := db.Model(item).Find(list)
	err := res.Error
	if err != nil {
		return err
	}

	// page
	if page != nil {
		count := res.RowsAffected
		page.Page = 1
		page.Size = int(count)
		page.Total = count
	}

	return nil
}

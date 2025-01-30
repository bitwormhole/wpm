package idao

import (
	"fmt"

	"github.com/starter-go/rbac"
	"gorm.io/gorm"
)

type finder struct {
	db    *gorm.DB
	page  *rbac.Pagination
	item  any
	list  any
	all   bool
	where func(*gorm.DB) *gorm.DB
}

func (inst *finder) checkParams() error {

	page := inst.page
	db := inst.db
	list := inst.list
	item := inst.item

	if item == nil {
		return fmt.Errorf("param: item is nil")
	}
	if list == nil {
		return fmt.Errorf("param: list is nil")
	}
	if db == nil {
		return fmt.Errorf("param: db is nil")
	}

	if page == nil {
		page = &rbac.Pagination{
			Page: 1,
			Size: 5,
		}
		inst.page = page
	}

	return nil
}

func (inst *finder) find() error {

	err := inst.checkParams()
	if err != nil {
		return err
	}
	page := inst.page
	db := inst.db
	list := inst.list
	item := inst.item
	all := inst.all
	where := inst.where

	// model
	db = db.Model(item)

	// where
	if where != nil {
		db = where(db)
	}

	// count
	count := int64(0)
	res := db.Count(&count)
	if res.Error == nil {
		page.Total = count
	}

	// page & size
	if !all {
		limit := page.Limit()
		offset := page.Offset()
		db = db.Limit(int(limit)).Offset(int(offset))
	}

	// find
	res = db.Find(&list)
	return res.Error
}

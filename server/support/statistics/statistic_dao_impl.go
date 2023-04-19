package statistics

import (
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/data/dao"
	"github.com/bitwormhole/wpm/server/data/dbagent"
	"github.com/bitwormhole/wpm/server/data/entity"
)

// ImpStatisticDao ...
type ImpStatisticDao struct {
	markup.Component `id:"StatisticDAO"`

	Agent dbagent.GormDBAgent `inject:"#GormDBAgent"`
}

func (inst *ImpStatisticDao) _Impl() dao.StatisticDAO {
	return inst
}

// Count ...
func (inst *ImpStatisticDao) count(m any) int64 {
	n := int64(0)
	db := inst.Agent.DB()
	db.Model(m).Count(&n)
	return n
}

// CountContentTypes ...
func (inst *ImpStatisticDao) CountContentTypes() int64 {
	return inst.count(&entity.ContentType{})
}

// CountIntentTemplates ...
func (inst *ImpStatisticDao) CountIntentTemplates() int64 {
	return inst.count(&entity.IntentTemplate{})
}

// CountRepositories ...
func (inst *ImpStatisticDao) CountRepositories() int64 {
	return inst.count(&entity.LocalRepository{})
}

// CountRemotes ...
func (inst *ImpStatisticDao) CountRemotes() int64 {
	// return inst.count(&entity.LocalRepository{})
	return 0
}

func (inst *ImpStatisticDao) CountBackups() int64 {
	return 0
}

func (inst *ImpStatisticDao) CountExecutables() int64 {
	return inst.count(&entity.Executable{})
}

func (inst *ImpStatisticDao) CountMediae() int64 {
	return inst.count(&entity.Media{})
}

func (inst *ImpStatisticDao) CountPlugIns() int64 {
	return inst.count(&entity.SoftwarePackage{})
}

func (inst *ImpStatisticDao) CountProjects() int64 {
	return inst.count(&entity.Project{})
}

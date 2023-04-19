package dao

// StatisticDAO ...
type StatisticDAO interface {
	CountBackups() int64

	CountContentTypes() int64

	CountExecutables() int64

	CountIntentTemplates() int64

	CountMediae() int64

	CountPlugIns() int64

	CountProjects() int64

	CountRepositories() int64

	CountRemotes() int64
}

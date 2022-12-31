package entity

// 表名前缀
const (
	TableNamePrefix = "wpm_"
)

// TableName ...
func (Executable) TableName() string {
	return TableNamePrefix + "executable"
}

// TableName ...
func (Intent) TableName() string {
	return TableNamePrefix + "intent"
}

// TableName ...
func (Repository) TableName() string {
	return TableNamePrefix + "repository"
}

// TableName ...
func (Project) TableName() string {
	return TableNamePrefix + "project"
}

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
func (RemoteRepository) TableName() string {
	return TableNamePrefix + "remote_repository"
}

// TableName ...
func (LocalRepository) TableName() string {
	return TableNamePrefix + "local_repository"
}

// TableName ...
func (MainRepository) TableName() string {
	return TableNamePrefix + "main_repository"
}

// TableName ...
func (Project) TableName() string {
	return TableNamePrefix + "project"
}

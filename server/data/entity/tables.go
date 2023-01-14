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

// TableName ...
func (User) TableName() string {
	return TableNamePrefix + "user"
}

////////////////////////////////////////////////////////////////////////////////

// ListPrototypes 罗列出所有的 entity 原型
func ListPrototypes() []any {
	list := make([]any, 0)

	list = append(list, &Executable{})
	list = append(list, &Intent{})
	list = append(list, &Project{})
	list = append(list, &LocalRepository{})
	list = append(list, &MainRepository{})
	list = append(list, &RemoteRepository{})
	list = append(list, &User{})

	return list
}

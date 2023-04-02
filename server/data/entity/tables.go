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
func (IntentTemplate) TableName() string {
	return TableNamePrefix + "intent_template"
}

// TableName ...
func (Media) TableName() string {
	return TableNamePrefix + "media"
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
func (Location) TableName() string {
	return TableNamePrefix + "location"
}

// TableName ...
func (Project) TableName() string {
	return TableNamePrefix + "project"
}

// TableName ...
func (ContentType) TableName() string {
	return TableNamePrefix + "content_type"
}

// TableName ...
func (Setting) TableName() string {
	return TableNamePrefix + "setting"
}

// TableName ...
func (User) TableName() string {
	return TableNamePrefix + "user"
}

// TableName ...
func (Worktree) TableName() string {
	return TableNamePrefix + "worktree"
}

////////////////////////////////////////////////////////////////////////////////

// ListPrototypes 罗列出所有的 entity 原型
func ListPrototypes() []any {
	list := make([]any, 0)

	list = append(list, &Executable{})
	list = append(list, &IntentTemplate{})
	list = append(list, &Project{})
	list = append(list, &ContentType{})
	list = append(list, &LocalRepository{})
	list = append(list, &Media{})
	list = append(list, &Location{})
	list = append(list, &RemoteRepository{})
	list = append(list, &Setting{})
	list = append(list, &User{})
	list = append(list, &Worktree{})

	return list
}

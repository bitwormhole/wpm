package service

// AppDataService ...
type AppDataService interface {
	GetAppDataDirectory() string

	GetSQLiteDBFile() string

	GetMainRepositoryPath() string
}

package wpm

import "github.com/starter-go/afs"

// DataDir 提供 ~/.wpm 文件夹下的各种路径
type DataDir interface {
	Path() afs.Path

	// 取 ~/.wpm 文件夹下的子目录或文件
	GetPath(path string) afs.Path

	// 取 ~/.wpm/files 文件夹
	Files() afs.Path

	// 取 ~/.wpm/files/:name 文件夹
	GetMediaBucket(name string) afs.Path
}

// Environment 提供 WPM 的运行环境信息
type Environment interface {
	BaseServerPort() int

	CurrentServerPort() int

	CurrentUserName() string

	CurrentUserHome() afs.Path

	CurrentExecutableFile() afs.Path

	DataDir() DataDir

	UseHTTPS() bool
}

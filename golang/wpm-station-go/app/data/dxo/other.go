package dxo

// StringList 表示一个包含多个词组的字符串，各个元素之间以逗号分隔
type StringList string

// FileState 表示文件或者目录的状态
type FileState string

// 定义文件或者目录的状态
const (
	FileStateInit      FileState = ""
	FileStateUntracked FileState = "untracked"
	FileStateReady     FileState = "ready"
	FileStateOffline   FileState = "offline"
)

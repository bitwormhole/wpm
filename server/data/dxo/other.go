package dxo

// FileState 表示文件或者目录的状态
type FileState string

// 定义文件或者目录的状态
const (
	FileStateUndefine    FileState = ""
	FileStateInit        FileState = "init"
	FileStateUntracked   FileState = "untracked"
	FileStateReady       FileState = "ready"
	FileStateOffline     FileState = "offline"
	FileStateMoved       FileState = "moved"
	FileStateUnknowError FileState = "error"
)

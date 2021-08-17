package base

// VO 所有VO的基本结构
type VO struct {
	Status    int
	Message   string
	Error     string
	Debug     string
	Timestamp int64
}

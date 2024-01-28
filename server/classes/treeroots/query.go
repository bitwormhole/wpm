package treeroots

import "github.com/starter-go/rbac"

// Query ...
type Query struct {
	All        bool // 查询全部条目
	Pagination rbac.Pagination
}

package namespaces

import "github.com/starter-go/rbac"

// Query 表示该类型的查询参数
type Query struct {
	All        bool // 查询全部条目
	Pagination rbac.Pagination
}

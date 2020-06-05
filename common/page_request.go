package common

// 分页请求参数结构体
type PageRequest struct {
	KeyWord string
	CurrentPage int
	PageSize int
}
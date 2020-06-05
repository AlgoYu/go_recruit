package common

type PageResult struct {
	Total int64 `json:"total"`
	PageSize int `json:"page_size"`
	Data interface{} `json:"data"`
}
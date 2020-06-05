package models

// 管理账号结构体
type Admin struct {
	Name string `orm:"pk" json:"name"`
	Picture string `json:"picture"`
	Password string `json:"password"`
	Introduce string `json:"introduce"`
}
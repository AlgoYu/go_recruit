package models

// 地区结构体
type Province struct {
	Id uint	`orm:"pk"`
	Name string
}
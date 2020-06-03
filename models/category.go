package models

// 理文科结构体
type Category struct {
	Id uint `orm:"pk"`
	Name string
}
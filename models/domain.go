package models

// 领域结构体
type Domain struct {
	Id uint `orm:"pk"`
	Name string
}
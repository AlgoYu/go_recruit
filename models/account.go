package models

// 账号结构体
type Account struct {
	Id string `orm:"pk"`
	Picture string
	Name string
	Introduce string
	Ic string
	Sex uint8
	Age uint8
	Email string
	CreateDatetime string
	UpdateDatetime string
}
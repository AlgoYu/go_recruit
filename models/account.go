package models

import "time"

// 账号结构体
type Account struct {
	Id uint `orm:"pk" json:"id"`
	Password string `json:"password"`
	Picture string `json:"picture"`
	Name string	`json:"name" valid:"Required"`
	Introduce string `json:"introduce"`
	Address string `json:"address"`
	Ic string `json:"ic" valid:"Required"`
	Contact string `json:"contact" valid:"Required"`
	Sex uint8 `json:"sex" valid:"Required"`
	Age uint8 `json:"age"`
	Email string `json:"email" valid:"Required"`
	CreateDatetime time.Time `json:"createDatetime"`
	UpdateDatetime time.Time `json:"updateDatetime"`
}
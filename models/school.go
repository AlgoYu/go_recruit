package models

import "time"

// 学校结构体
type School struct {
	Id uint	`orm:"pk"`
	Logo string
	Cover string
	Name string
	code string
	Address string
	Introduce string
	Contact string
	Latitude float32
	Longitude float32
	ProvinceId uint
	CreateDatetime time.Time
	UpdateDatetime time.Time
}
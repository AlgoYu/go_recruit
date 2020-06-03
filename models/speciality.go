package models

// 专业结构体
type Speciality struct {
	Id uint `orm:"pk"`
	Name string
	DomainId uint
	CategoryId uint
	Tuition float32
}
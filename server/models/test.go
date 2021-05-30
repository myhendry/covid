package models

type Test struct {
	ID uint
	Title string `json:"title"`
	Quantity uint `json:"quantity"`
	Approved bool `json:"approved"`
	Active bool `json:"active" gorm:"default:true"`
}
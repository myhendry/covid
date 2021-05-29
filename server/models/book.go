package models

type Book struct {
	ID uint
	Title string `json:"title"`
	UserID uint `json:"user_id"`
}
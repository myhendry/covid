package models

type Book struct {
	ID uint
	Title string `json:"title"`
	UserID uint `json:"user_id"`
	Authors []Author `json:"authors" gorm:"many2many:books_authors"`
}

type Author struct {
	ID uint
	Name string `json:"name"`
}
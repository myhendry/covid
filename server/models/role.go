package models

// User : Role - One 2 One Association
type Role struct {
	ID uint
	Name string `json:"name"`
}
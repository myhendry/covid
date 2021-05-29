package models

// User : Profile - One 2 One Association. Create User then Create Profile
type Profile struct {
	ID uint
	Country string `json:"country"`
	UserID uint `json:"user_id"`
}
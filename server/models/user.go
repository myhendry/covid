package models

import "time"

type User struct {
	ID uint
	Name string `json:"name"`
	Age uint	`json:"age"`
	Profile Profile `gorm:"constraint:OnDelete:CASCADE"`
	RoleID uint `json:"role_id"`
	Role Role `gorm:"constraint:OnDelete:CASCADE"`
	Books []Book `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt    time.Time
  UpdatedAt    time.Time
}
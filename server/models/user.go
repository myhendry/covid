package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Age uint	`json:"age"`
	Email string `json:"email"`
	Password []byte `json:"password"`
	Profile Profile `gorm:"constraint:OnDelete:CASCADE"`
	RoleID uint `json:"role_id"`
	Role Role `json:"role" gorm:"foreignKey:RoleID"`
	Books []Book `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt    time.Time
  UpdatedAt    time.Time
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
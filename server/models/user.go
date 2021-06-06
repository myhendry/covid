package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Age uint	`json:"age"`
	Nickname string `json:"nickname" gorm:"-"`
	Email string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
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

func (user *User) GetNickname() string {
	return user.Name + " " + user.Email;
}
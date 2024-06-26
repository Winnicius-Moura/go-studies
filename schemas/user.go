package schemas

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Username string `gorm:"unique"`
	Password string
}

type UserResponse struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Token     string `json:"token"`
}

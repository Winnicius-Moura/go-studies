package schemas

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Username string `gorm:"unique"`
	Password string
	Profile  string `gorm:"type:varchar(20);not null"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Profile   string `json:"profile"`
	Token     string `json:"token"`
}

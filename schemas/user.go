package schemas

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
}

type UserResponse struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Username  string `json:"username"`
	Token  string `json:"token"`
}

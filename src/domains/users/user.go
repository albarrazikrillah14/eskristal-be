package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"primaryKey;column:id"`
	RoleID    string    `gorm:"column:role_id"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email;unique"`
	Username  string    `gorm:"column:username;unique"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-create"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

func (u *User) TableName() string {
	return "users"
}

type CreateUserRequest struct {
	RoleID   string `json:"role_id" validate:"required,uuid"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,min=8"`
}

func (c *CreateUserRequest) MapToUser() User {
	return User{
		Name:     c.Name,
		Email:    c.Email,
		Username: c.Username,
		Password: c.Password,
		RoleID:   c.RoleID,
	}
}

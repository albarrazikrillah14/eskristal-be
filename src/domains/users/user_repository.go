package users

import (
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, tx *gorm.DB, user *User) error
	VerifyUsernameIsNotExists(ctx context.Context, tx *gorm.DB, username string) error
	VerifyEmailIsNotExists(ctx context.Context, tx *gorm.DB, email string) error
}

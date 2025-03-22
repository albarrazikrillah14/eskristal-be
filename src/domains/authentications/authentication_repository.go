package authentications

import (
	"context"

	"gorm.io/gorm"
)

type AuthenticationRepository interface {
	Create(ctx context.Context, tx *gorm.DB, token Authentication) error
	Delete(ctx context.Context, tx *gorm.DB, token string) error
	FindByToken(ctx context.Context, tx *gorm.DB, token string) (*Authentication, error)
}

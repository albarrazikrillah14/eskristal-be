package roles

import (
	"context"

	"gorm.io/gorm"
)

type RoleRepository interface {
	FindByID(ctx context.Context, tx *gorm.DB, id string) (*Role, error)
	VerifyRoleIsNotExists(ctx context.Context, tx *gorm.DB, name string) error
	Create(ctx context.Context, tx *gorm.DB, role *Role) error
	FindAll(ctx context.Context) []Role
}

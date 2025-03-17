package roles

import (
	"context"

	"gorm.io/gorm"
)

type RoleRepository interface {
	FindByID(ctx context.Context, tx *gorm.DB, id string) (Role, error)
}

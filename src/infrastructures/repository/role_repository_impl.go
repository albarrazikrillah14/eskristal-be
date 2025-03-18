package repository

import (
	"context"
	idgenerator "rania-eskristal/src/applications/id_generator"
	"rania-eskristal/src/commons/exceptions"
	"rania-eskristal/src/domains/roles"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type roleRepositoryImpl struct {
	DB          *gorm.DB
	Logger      *logrus.Logger
	IDGenerator idgenerator.IDGenerator
}

func NewRoleRepositoryImpl(
	db *gorm.DB,
	logger *logrus.Logger,
	idGenerator idgenerator.IDGenerator,
) roles.RoleRepository {
	return &roleRepositoryImpl{
		DB:          db,
		Logger:      logger,
		IDGenerator: idGenerator,
	}
}

func (r *roleRepositoryImpl) FindByID(ctx context.Context, tx *gorm.DB, id string) (*roles.Role, error) {
	traceID := ctx.Value("trace_id")
	role := roles.Role{}

	result := tx.Select("id", "name").Take(&role, "id = ?", id)
	if result.Error != nil {
		r.Logger.Error(exceptions.NewLogBody(
			traceID,
			result.Error,
		))

		return nil, exceptions.NewInvariantError("ERR_UNKNOWN")
	}

	if result.RowsAffected == 0 {
		r.Logger.Info(
			traceID,
			"ERR_ROLE_NOT_FOUND",
		)

		return nil, exceptions.NewNotFoundError("ERR_ROLE_NOT_FOUND")
	}

	return &role, nil
}

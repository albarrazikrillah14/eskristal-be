package repository

import (
	"context"
	"errors"
	idgenerator "rania-eskristal/src/applications/id_generator"
	"rania-eskristal/src/commons/enums"
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
	traceID := ctx.Value(enums.TraceIDKey)
	role := roles.Role{}

	result := tx.Select("id", "name").Take(&role, "id = ?", id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			r.Logger.WithFields(logrus.Fields{
				enums.TraceIDKey: traceID,
				enums.ErrorsKey:  "ERR_ROLE.NOT_FOUND",
			}).Error("ERR_ROLE.NOT_FOUND")

			return nil, exceptions.NewNotFoundError("ERR_ROLE.NOT_FOUND")
		}

		r.Logger.WithFields(logrus.Fields{
			enums.TraceIDKey: traceID,
			enums.ErrorsKey:  result.Error.Error(),
		}).Error("ERR_UNKNOWN")

		return nil, exceptions.NewInvariantError("ERR_UNKNOWN")
	}

	return &role, nil
}

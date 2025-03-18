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

// Create implements roles.RoleRepository.
func (r *roleRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, role *roles.Role) error {
	traceID := ctx.Value(enums.TraceIDKey)

	role.ID = r.IDGenerator.Generate()

	result := tx.Create(role)

	if result.Error != nil {
		r.Logger.WithFields(logrus.Fields{
			enums.TraceIDKey: traceID,
			enums.ErrorsKey:  result.Error.Error(),
		}).Error("ERR_UKNOWN")

		return exceptions.NewInvariantError("ERR_ROLE_UNKNOWN")
	}

	return nil
}

// VerifyRoleIsNotExists implements roles.RoleRepository.
func (r *roleRepositoryImpl) VerifyRoleIsNotExists(ctx context.Context, tx *gorm.DB, name string) error {
	traceID := ctx.Value(enums.TraceIDKey)

	role := roles.Role{}

	result := tx.Select("id", "name").Take(&role, "name = ?", name)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}

		r.Logger.WithFields(logrus.Fields{
			enums.TraceIDKey: traceID,
			enums.ErrorsKey:  result.Error.Error(),
		}).Error("ERR_UKNOWN")

		return exceptions.NewInvariantError("ERR_ROLE_UNKNOWN")
	}

	r.Logger.WithFields(logrus.Fields{
		enums.TraceIDKey: traceID,
		enums.ErrorsKey:  "ERR_ROLE.NAME_DUPLICATE_KEY",
	}).Error("ERR_DUPLICATE_KEY")

	return exceptions.NewInvariantError("ERR_ROLE.NAME_DUPLICATE_KEY")
}

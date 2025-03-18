package repository

import (
	"context"
	"errors"
	idgenerator "rania-eskristal/src/applications/id_generator"
	"rania-eskristal/src/commons/enums"
	"rania-eskristal/src/commons/exceptions"
	"rania-eskristal/src/domains/users"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	DB          *gorm.DB
	Logger      *logrus.Logger
	IDGenerator idgenerator.IDGenerator
}

func NewUserRepositoryImpl(
	db *gorm.DB,
	logger *logrus.Logger,
	idGenerator idgenerator.IDGenerator,
) users.UserRepository {
	return &userRepositoryImpl{
		DB:          db,
		Logger:      logger,
		IDGenerator: idGenerator,
	}
}

// Create implements users.UserRepository.
func (u *userRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, user *users.User) error {
	traceID := ctx.Value(enums.TraceIDKey)
	user.ID = u.IDGenerator.Generate()

	result := tx.Create(user)

	if result.Error != nil {
		u.Logger.WithFields(logrus.Fields{
			enums.TraceIDKey: traceID,
			enums.ErrorsKey:  result.Error.Error(),
		}).Error("ERR_CREATE_USER")

		return exceptions.NewInvariantError("ERR_CREATE_USER")
	}

	return nil
}

// VerifyEmailIsExists implements users.UserRepository.
func (u *userRepositoryImpl) VerifyEmailIsNotExists(ctx context.Context, tx *gorm.DB, email string) error {
	traceID := ctx.Value(enums.TraceIDKey)

	user := users.User{}

	result := u.DB.Select("email").Take(&user, "email = ?", email)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}

		u.Logger.WithFields(logrus.Fields{
			enums.TraceIDKey: traceID,
			enums.ErrorsKey:  result.Error.Error(),
		}).Error("ERR_UNKNOWN")

		return exceptions.NewInvariantError("ERR_UNKNOWN")
	}

	if result.RowsAffected == 0 {
		return nil
	}

	u.Logger.WithFields(logrus.Fields{
		enums.TraceIDKey: traceID,
		enums.ErrorsKey:  "ERR_USER.EMAIL_DUPLICATE_KEY",
	}).Error("ERR_DUPLICATE_KEY")

	return exceptions.NewInvariantError("ERR_USER.EMAIL_DUPLICATE_KEY")
}

// VerifyUsernameIsExists implements users.UserRepository.
func (u *userRepositoryImpl) VerifyUsernameIsNotExists(ctx context.Context, tx *gorm.DB, username string) error {
	traceID := ctx.Value(enums.TraceIDKey)

	user := users.User{}

	result := u.DB.Select("username").Take(&user, "email = ?", username)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}

		u.Logger.WithFields(logrus.Fields{
			enums.TraceIDKey: traceID,
			enums.ErrorsKey:  result.Error.Error(),
		}).Error("ERR_UNKNOWN")

		return exceptions.NewInvariantError("ERR_UNKNOWN")
	}

	if result.RowsAffected == 0 {
		return nil
	}

	u.Logger.WithFields(logrus.Fields{
		enums.TraceIDKey: traceID,
		enums.ErrorsKey:  "ERR_USER.USERNAME_DUPLICATE_KEY",
	}).Error("ERR_DUPLICATE_KEY")

	return exceptions.NewInvariantError("ERR_USER.USERNAME_DUPLICATE_KEY")
}

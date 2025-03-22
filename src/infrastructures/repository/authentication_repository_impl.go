package repository

import (
	"context"
	"errors"
	"rania-eskristal/src/commons/enums"
	"rania-eskristal/src/commons/exceptions"
	"rania-eskristal/src/domains/authentications"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type authenticationRepositoryImpl struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewAuthenticationRepositoryImpl(
	db *gorm.DB,
	logger *logrus.Logger,
) authentications.AuthenticationRepository {
	return &authenticationRepositoryImpl{
		DB:     db,
		Logger: logger,
	}
}

// Create implements authentications.AuthenticationRepository.
func (a *authenticationRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, token authentications.Authentication) error {
	traceID := ctx.Value("trace_id")

	result := tx.Save(&token)

	if result.Error != nil {
		a.Logger.WithFields(logrus.Fields{
			enums.TraceIDKey: traceID,
			enums.ErrorsKey:  result.Error,
		})

		return exceptions.NewInvariantError("ERR_UNKNOWN")
	}

	return nil
}

// Delete implements authentications.AuthenticationRepository.
func (a *authenticationRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, token string) error {
	traceID := ctx.Value("trace_id")

	result := tx.Delete(&authentications.Authentication{}, "token = ?", token)

	if result.Error != nil {
		a.Logger.WithFields(logrus.Fields{
			enums.TraceIDKey: traceID,
			enums.ErrorsKey:  result.Error,
		})

		return exceptions.NewInvariantError("ERR_UNKNOWN")
	}

	return nil
}

// FindByToken implements authentications.AuthenticationRepository.
func (a *authenticationRepositoryImpl) FindByToken(ctx context.Context, tx *gorm.DB, token string) (*authentications.Authentication, error) {
	traceID := ctx.Value("trace_id")

	auth := authentications.Authentication{}

	result := tx.Find(&auth, "token = ?", token)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, exceptions.NewNotFoundError("ERR_AUTHENTICATION.TOKEN_NOT_FOUND")
		}

		a.Logger.WithFields(logrus.Fields{
			enums.TraceIDKey: traceID,
			enums.ErrorsKey:  result.Error,
		})

		return nil, exceptions.NewInvariantError("ERR_UNKNOWN")
	}

	return &auth, nil
}

package repository

import (
	"context"
	"rania-eskristal/src/commons/exceptions"
	"rania-eskristal/src/domains/users"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	DB            *gorm.DB
	Logger        *logrus.Logger
	UUIDGenerator func() string
}

func NewUserRepositoryImpl(
	db *gorm.DB,
	logger *logrus.Logger,
	uuidGenerator func() string,
) users.UserRepository {
	return &userRepositoryImpl{
		DB:            db,
		Logger:        logger,
		UUIDGenerator: uuidGenerator,
	}
}

// Create implements users.UserRepository.
func (u *userRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, user *users.User) error {
	traceID := ctx.Value("trace_id")
	user.ID = u.UUIDGenerator()

	result := tx.Create(user)

	if result.Error != nil {
		u.Logger.Error(exceptions.NewLogBody(
			traceID,
			result.Error,
		))

		return exceptions.NewInvariantError("ERR_CREATE_USER")
	}

	if result.RowsAffected == 0 {
		u.Logger.Error(exceptions.NewLogBody(
			traceID,
			"user cannot be inserted into database",
		))
	}

	return nil
}

// VerifyEmailIsExists implements users.UserRepository.
func (u *userRepositoryImpl) VerifyEmailIsNotExists(ctx context.Context, tx *gorm.DB, email string) error {
	traceID := ctx.Value("trace_id")

	user := users.User{}

	result := u.DB.Select("email").Take(&user, "email = ?", email)

	if result.Error != nil {
		u.Logger.Error(exceptions.NewLogBody(
			traceID,
			result.Error,
		))

		return exceptions.NewInvariantError("ERR_UNKNOWN")
	}

	if result.RowsAffected == 0 {
		return nil
	}

	u.Logger.Info(exceptions.NewLogBody(
		traceID,
		"ERR_USER.EMAIL_DUPLICATE_KEY",
	))

	return exceptions.NewInvariantError("ERR_USER.EMAIL_DUPLICATE_KEY")
}

// VerifyUsernameIsExists implements users.UserRepository.
func (u *userRepositoryImpl) VerifyUsernameIsNotExists(ctx context.Context, tx *gorm.DB, username string) error {
	traceID := ctx.Value("trace_id")

	user := users.User{}

	result := u.DB.Select("username").Take(&user, "username = ?", username)

	if result.Error != nil {
		u.Logger.Error(exceptions.NewLogBody(
			traceID,
			result.Error,
		))

		return exceptions.NewInvariantError("ERR_UNKNOWN")
	}

	if result.RowsAffected == 0 {
		return nil
	}

	u.Logger.Info(exceptions.NewLogBody(
		traceID,
		"ERR_USER.USERNMAE_DUPLICATE_KEY",
	))

	return exceptions.NewInvariantError("ERR_USER.USERNMAE_DUPLICATE_KEY")
}

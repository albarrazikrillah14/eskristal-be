package usecase

import (
	"context"
	"rania-eskristal/src/applications/security"
	"rania-eskristal/src/commons/exceptions"
	"rania-eskristal/src/commons/helper"
	"rania-eskristal/src/domains/roles"
	"rania-eskristal/src/domains/users"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserUseCase interface {
	Create(ctx context.Context, request *users.CreateUserRequest) error
}

type userUseCaseImpl struct {
	DB             *gorm.DB
	Validate       *validator.Validate
	Logger         *logrus.Logger
	Hash           security.Hash
	RoleRepository roles.RoleRepository
	UserRepository users.UserRepository
}

func NewUserUseCaseImpl(
	db *gorm.DB,
	validator *validator.Validate,
	logger *logrus.Logger,
	hash security.Hash,
	roleRepository roles.RoleRepository,
	userRepository users.UserRepository,
) UserUseCase {
	return &userUseCaseImpl{
		DB:             db,
		Validate:       validator,
		Logger:         logger,
		Hash:           hash,
		RoleRepository: roleRepository,
		UserRepository: userRepository,
	}
}

func (u *userUseCaseImpl) Create(ctx context.Context, request *users.CreateUserRequest) error {

	traceID := ctx.Value("trace_id")
	u.Logger.Info(exceptions.NewLogBody(
		traceID,
		request,
	))

	err := helper.NewValidationStruct(u.Validate, request)
	if err != nil {
		return err
	}

	errTx := u.DB.Transaction(func(tx *gorm.DB) error {
		_, err := u.RoleRepository.FindByID(ctx, tx, request.RoleID)

		if err != nil {
			return err
		}

		err = u.UserRepository.VerifyUsernameIsNotExists(ctx, tx, request.Username)
		if err != nil {
			return err
		}

		err = u.UserRepository.VerifyEmailIsNotExists(ctx, tx, request.Email)
		if err != nil {
			return err
		}

		user := request.MapToUser()
		user.Password = u.Hash.Hash(user.Password)

		err = u.UserRepository.Create(ctx, tx, &user)

		return err
	})

	return errTx
}

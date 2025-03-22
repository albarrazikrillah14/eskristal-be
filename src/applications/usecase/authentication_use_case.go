package usecase

import (
	"context"
	"rania-eskristal/src/applications/security"
	"rania-eskristal/src/commons/enums"
	"rania-eskristal/src/commons/helpers"
	"rania-eskristal/src/domains/authentications"
	"rania-eskristal/src/domains/users"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthenticationUseCase interface {
	Create(ctx context.Context, request *authentications.AuthenticationRequest) (*string, error)
}

type authenticationUseCaseImpl struct {
	DB               *gorm.DB
	Validator        *validator.Validate
	Logger           *logrus.Logger
	AuthTokenManager security.AuthenticationTokenManager
	Hash             security.Hash
	UserRepository   users.UserRepository
	// AuthenticationRepository authentications.AuthenticationRepository
}

func NewAuthenticationUseCaseImpl(
	db *gorm.DB,
	validator *validator.Validate,
	logger *logrus.Logger,
	authTokenManager security.AuthenticationTokenManager,
	hash security.Hash,
	userRepository users.UserRepository,
	// authenticationRepository authentications.AuthenticationRepository,
) AuthenticationUseCase {
	return &authenticationUseCaseImpl{
		DB:               db,
		Validator:        validator,
		Logger:           logger,
		AuthTokenManager: authTokenManager,
		Hash:             hash,
		UserRepository:   userRepository,
		// AuthenticationRepository: authenticationRepository,
	}
}

// Upsert implements AuthenticationUseCase.
func (a *authenticationUseCaseImpl) Create(ctx context.Context, request *authentications.AuthenticationRequest) (*string, error) {
	traceID := ctx.Value(enums.TraceIDKey)

	token := ""
	errTx := a.DB.Transaction(func(tx *gorm.DB) error {

		err := helpers.NewValidationStruct(a.Validator, request, a.Logger, traceID)
		if err != nil {
			return err
		}

		user, err := a.UserRepository.FindByEmailOrUsername(ctx, tx, request.UsernameOrEmail)
		if err != nil {
			return err
		}

		err = a.Hash.Compare(user.Password, request.Password)
		if err != nil {
			return err
		}

		accessToken, err := a.AuthTokenManager.Generate(ctx, &authentications.AuthenticationPayload{UserID: user.ID, RoleID: user.RoleID})
		if err != nil {
			return err
		}

		token = *accessToken

		return nil
	})

	if errTx != nil {
		return nil, errTx
	}

	return &token, nil
}

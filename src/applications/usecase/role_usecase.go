package usecase

import (
	"context"
	"rania-eskristal/src/commons/enums"
	"rania-eskristal/src/commons/helpers"
	"rania-eskristal/src/domains/roles"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RoleUseCase interface {
	Create(ctx context.Context, request *roles.CreateRoleRequest) error
	FindAll(ctx context.Context) []roles.RoleResponse
}

type roleUseCaseImpl struct {
	DB             *gorm.DB
	Validate       *validator.Validate
	Logger         *logrus.Logger
	RoleRepository roles.RoleRepository
}

func NewRoleUseCaseImpl(
	db *gorm.DB,
	validator *validator.Validate,
	logger *logrus.Logger,
	roleRepository roles.RoleRepository,
) RoleUseCase {
	return &roleUseCaseImpl{
		DB:             db,
		Validate:       validator,
		Logger:         logger,
		RoleRepository: roleRepository,
	}
}

// Create implements UserUseCase.
func (r *roleUseCaseImpl) Create(ctx context.Context, request *roles.CreateRoleRequest) error {
	traceID := ctx.Value(enums.TraceIDKey)
	r.Logger.WithFields(logrus.Fields{
		enums.TraceIDKey: traceID,
		enums.PayloadKey: *request,
	}).Info("ROLE_USECASE.CREATE_CALLED")

	err := helpers.NewValidationStruct(r.Validate, request, r.Logger, traceID)
	if err != nil {
		return err
	}

	errTx := r.DB.Transaction(func(tx *gorm.DB) error {
		err := r.RoleRepository.VerifyRoleIsNotExists(ctx, tx, request.Name)
		if err != nil {
			return err
		}

		role := request.MapToRole()
		err = r.RoleRepository.Create(ctx, tx, &role)

		return err
	})

	return errTx
}

// FindAll implements RoleUseCase.
func (r *roleUseCaseImpl) FindAll(ctx context.Context) []roles.RoleResponse {
	result := r.RoleRepository.FindAll(ctx)

	response := []roles.RoleResponse{}

	for _, rl := range result {
		response = append(response, rl.MapToResponse())
	}

	return response
}

package handlers

import (
	"context"
	idgenerator "rania-eskristal/src/applications/id_generator"
	"rania-eskristal/src/applications/usecase"
	"rania-eskristal/src/commons/enums"
	"rania-eskristal/src/commons/exceptions"
	"rania-eskristal/src/domains/users"
	"rania-eskristal/src/domains/web"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	PostUserHandler(ctx *fiber.Ctx) error
}

type UserHandlerImpl struct {
	UseCase     usecase.UserUseCase
	IDGenerator idgenerator.IDGenerator
}

func NewUserHandlerImpl(
	usecase usecase.UserUseCase,
	idGenerator idgenerator.IDGenerator,
) UserHandler {
	return &UserHandlerImpl{
		UseCase:     usecase,
		IDGenerator: idGenerator,
	}
}

func (u *UserHandlerImpl) PostUserHandler(ctx *fiber.Ctx) error {
	request := users.CreateUserRequest{}
	err := ctx.BodyParser(&request)

	if err != nil {
		return exceptions.NewInvariantError(err.Error())
	}

	c := context.Background()
	traceID := u.IDGenerator.Generate()
	ctx.Locals(enums.TraceIDKey, traceID)
	contextTrace := context.WithValue(c, enums.TraceIDKey, traceID)

	err = u.UseCase.Create(contextTrace, &request)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(web.NewBaseResponse(
		traceID,
		"berhasil menambahkan user",
	))
}

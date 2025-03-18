package handlers

import (
	"context"
	idgenerator "rania-eskristal/src/applications/id_generator"
	"rania-eskristal/src/applications/usecase"
	"rania-eskristal/src/commons/enums"
	"rania-eskristal/src/commons/exceptions"
	"rania-eskristal/src/domains/roles"
	"rania-eskristal/src/domains/web"

	"github.com/gofiber/fiber/v2"
)

type RoleHandler interface {
	PostRoleHandler(ctx *fiber.Ctx) error
}

type RoleHandlerImpl struct {
	UseCase     usecase.RoleUseCase
	IDGenerator idgenerator.IDGenerator
}

func NewRoleHandlerImpl(
	usecase usecase.RoleUseCase,
	idGenerator idgenerator.IDGenerator,
) RoleHandler {
	return &RoleHandlerImpl{
		UseCase:     usecase,
		IDGenerator: idGenerator,
	}
}

// PostRoleHandler implements RoleHandler.
func (r *RoleHandlerImpl) PostRoleHandler(ctx *fiber.Ctx) error {
	c := context.Background()
	traceID := r.IDGenerator.Generate()
	ctx.Locals(enums.TraceIDKey, traceID)
	contextTrace := context.WithValue(c, enums.TraceIDKey, traceID)

	request := roles.CreateRoleRequest{}
	err := ctx.BodyParser(&request)

	if err != nil {
		return exceptions.NewInvariantError(err.Error())
	}

	err = r.UseCase.Create(contextTrace, &request)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(web.NewBaseResponse(
		traceID,
		"berhasil menambahkan role",
	))
}

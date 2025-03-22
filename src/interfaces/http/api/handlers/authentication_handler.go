package handlers

import (
	"context"
	idgenerator "rania-eskristal/src/applications/id_generator"
	"rania-eskristal/src/applications/usecase"
	"rania-eskristal/src/domains/authentications"
	"rania-eskristal/src/domains/web"

	"github.com/gofiber/fiber/v2"
)

type AuthenticationHandler interface {
	PostAuthenticationHandler(ctx *fiber.Ctx) error
	DeleteAuthenticationByIDHandler(ctx *fiber.Ctx) error
}

type AuthenticationHandlerImpl struct {
	UseCase     usecase.AuthenticationUseCase
	IDGenerator idgenerator.IDGenerator
}

func NewAuthenticationHandlerImpl(
	usecase usecase.AuthenticationUseCase,
	idGenerator idgenerator.IDGenerator,
) AuthenticationHandler {
	return &AuthenticationHandlerImpl{
		UseCase:     usecase,
		IDGenerator: idGenerator,
	}
}

// DeleteAuthenticationByIDHandler implements AuthenticationHandler.
func (a *AuthenticationHandlerImpl) DeleteAuthenticationByIDHandler(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// PostAuthenticationHandler implements AuthenticationHandler.
func (a *AuthenticationHandlerImpl) PostAuthenticationHandler(ctx *fiber.Ctx) error {
	request := authentications.AuthenticationRequest{}

	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	accessToken, err := a.UseCase.Create(context.Background(), &request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(web.NewBaseResponse(
		nil,
		accessToken,
	))
}

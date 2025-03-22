package http

import (
	"rania-eskristal/src/applications/usecase"
	"rania-eskristal/src/commons/config"
	"rania-eskristal/src/infrastructures/database/pg"
	idgenerator "rania-eskristal/src/infrastructures/id_generator"
	"rania-eskristal/src/infrastructures/repository"
	"rania-eskristal/src/infrastructures/security"
	"rania-eskristal/src/interfaces/http/api/handlers"
	"rania-eskristal/src/interfaces/http/api/middlewares"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func New(config *config.Config, logger *logrus.Logger) *fiber.App {
	app := fiber.New(
		fiber.Config{
			ErrorHandler: middlewares.ErrorHandler,
		},
	)

	//external
	bcrypt := security.NewBcryptHash()
	db := pg.New(&config.DB, logger).Connection()
	validator := validator.New()
	idGenerator := idgenerator.New()
	jwtTokenManager := security.NewJwtTokenManager(&config.JWT)

	//repository
	userRepository := repository.NewUserRepositoryImpl(db, logger, idGenerator)
	roleRepository := repository.NewRoleRepositoryImpl(db, logger, idGenerator)

	//usecase
	userUseCase := usecase.NewUserUseCaseImpl(db, validator, logger, bcrypt, roleRepository, userRepository)
	roleUseCase := usecase.NewRoleUseCaseImpl(db, validator, logger, roleRepository)
	authenticationUseCase := usecase.NewAuthenticationUseCaseImpl(db, validator, logger, jwtTokenManager, bcrypt, userRepository)

	//handlers
	userHandler := handlers.NewUserHandlerImpl(userUseCase, idGenerator)
	roleHandler := handlers.NewRoleHandlerImpl(roleUseCase, idGenerator)
	authenticationHandler := handlers.NewAuthenticationHandlerImpl(authenticationUseCase, idGenerator)

	//users
	app.Post("/users", userHandler.PostUserHandler)

	//roles
	app.Post("/roles", roleHandler.PostRoleHandler)
	app.Get("/roles", roleHandler.GetRolesHandler)
	app.Delete("/roles/:id", roleHandler.DeleteRoleByIDHandler)

	//authentication
	app.Post("/authentications", authenticationHandler.PostAuthenticationHandler)
	return app
}

package middlewares

import (
	"rania-eskristal/src/commons/enums"
	"rania-eskristal/src/commons/exceptions"
	"rania-eskristal/src/domains/web"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	traceID := ctx.Locals(enums.TraceIDKey)

	if err != nil {
		if clientError, ok := err.(*exceptions.ClientError); ok {
			if clientError.Name == "ValidationError" {
				return ctx.Status(clientError.StatusCode).JSON(web.NewBaseErrorResponse(
					traceID,
					clientError.ErrorValidation,
				))
			}
			return ctx.Status(clientError.StatusCode).JSON(web.NewBaseErrorResponse(
				traceID,
				clientError.ErrorMessage,
			))
		}

		if fiberErr, ok := err.(*fiber.Error); ok {
			return ctx.Status(fiberErr.Code).JSON(web.NewBaseErrorResponse(
				traceID,
				fiberErr.Error(),
			))
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(web.NewBaseErrorResponse(
			traceID,
			"Internal Server Error",
		))
	}

	return ctx.Next()

}

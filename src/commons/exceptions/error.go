package exceptions

import "github.com/gofiber/fiber/v2"

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ClientError struct {
	StatusCode      int
	Name            string
	ErrorMessage    string
	ErrorValidation []ValidationError
}

func (c *ClientError) Error() string {
	return c.ErrorMessage
}

func NewInvariantError(message string) error {
	return &ClientError{
		StatusCode:   fiber.StatusBadRequest,
		Name:         "InvariantError",
		ErrorMessage: message,
	}
}

func NewValidationError(err []ValidationError) error {
	return &ClientError{
		StatusCode:      fiber.StatusBadRequest,
		Name:            "ValidationError",
		ErrorValidation: err,
	}
}

func NewAuthenticationError(message string) error {
	return &ClientError{
		StatusCode:   fiber.StatusUnauthorized,
		Name:         "AuthenticationError",
		ErrorMessage: message,
	}
}

func NewNotFoundError(message string) error {
	return &ClientError{
		StatusCode:   fiber.StatusNotFound,
		Name:         "NotFoundError",
		ErrorMessage: message,
	}
}

func NewAuthorizationError(message string) error {
	return &ClientError{
		StatusCode:   fiber.StatusForbidden,
		Name:         "AuthorizationError",
		ErrorMessage: message,
	}
}

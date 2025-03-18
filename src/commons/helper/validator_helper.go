package helper

import (
	"fmt"
	"rania-eskristal/src/commons/exceptions"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func NewValidationStruct(
	val *validator.Validate,
	request any,
	logger *logrus.Logger,
	traceID any,
) error {
	err := val.Struct(request)
	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return fmt.Errorf("unexpected validation error type: %v", err)
	}

	errors := make([]exceptions.ValidationError, 0, len(validationErrors))
	// t := reflect.TypeOf(request).Elem()

	for _, fieldError := range validationErrors {
		// field, ok := t.FieldByName(fieldError.StructField())
		error := exceptions.ValidationError{
			Field:   fieldError.StructField(),
			Message: getValidationMessage(fieldError),
		}
		errors = append(errors, error)
	}

	logger.WithFields(logrus.Fields{
		"trace_id": traceID,
		"errors":   errors,
	}).Error("ERR_VALIDATION")

	return exceptions.NewValidationError(errors)

}

func getValidationMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "tidak boleh kosong"
	case "email":
		return "gunakan format email yang benar"
	case "min":
		return fmt.Sprintf("panjang karakter minimal adalah %s", fieldError.Param())
	case "max":
		return fmt.Sprintf("panjang karakter tidak boleh lebih dari %s", fieldError.Param())
	case "oneof":
		return fmt.Sprintf("gunakan salah satu dari : %v", fieldError.Param())
	case "datetime":
		return "format waktu tidak valid, gunakan format RFC3339: YYYY-MM-DDTHH:MM:SSZ"
	default:
		return fmt.Sprintf("Validation failed on %s with constraint: %s",
			fieldError.Field(), fieldError.Tag())
	}
}

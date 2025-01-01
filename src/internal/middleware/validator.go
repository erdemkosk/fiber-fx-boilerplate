package middleware

import (
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateStruct(payload interface{}) []*model.ValidationError {
	var errors []*model.ValidationError
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element model.ValidationError
			element.Field = err.Field()
			element.Message = getErrorMsg(err)
			errors = append(errors, &element)
		}
	}
	return errors
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Please enter a valid email address"
	case "min":
		return "Does not meet minimum length requirement"
	case "max":
		return "Exceeds maximum length limit"
	default:
		return "Validation error"
	}
}

func RequestValidator[T any](payload T) fiber.Handler {
	return func(c *fiber.Ctx) error {

		if err := c.BodyParser(&payload); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
		}

		if errors := ValidateStruct(payload); len(errors) > 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": errors,
			})
		}

		c.Locals("payload", payload)

		return c.Next()
	}
}

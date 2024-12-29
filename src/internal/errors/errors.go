package errors

import "github.com/gofiber/fiber/v2"

func NewNotFoundError(message string) *fiber.Error {
	return fiber.NewError(fiber.StatusNotFound, message)
}

func NewBadRequestError(message string) *fiber.Error {
	return fiber.NewError(fiber.StatusBadRequest, message)
}

func NewInternalError(message string) *fiber.Error {
	return fiber.NewError(fiber.StatusInternalServerError, message)
}

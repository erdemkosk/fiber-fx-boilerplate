package handler

import (
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/errors"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/model"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/service"
	"github.com/gofiber/fiber/v2"
)

type FooHandler struct {
	fooService *service.FooService
}

func NewFooHandler(fooService *service.FooService) *FooHandler {
	return &FooHandler{
		fooService: fooService,
	}
}

func (f *FooHandler) GetAll(c *fiber.Ctx) error {
	foos, err := f.fooService.GetAll()
	if err != nil {
		return errors.NewInternalError(err.Error())
	}
	return c.JSON(foos)
}

func (f *FooHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	foo, err := f.fooService.GetById(id)
	if err != nil {
		return errors.NewInternalError(err.Error())
	}
	if foo == nil {
		return errors.NewNotFoundError("Foo not found")
	}
	return c.JSON(foo)
}

func (f *FooHandler) Create(c *fiber.Ctx) error {
	foo := new(model.Foo)
	if err := c.BodyParser(foo); err != nil {
		return errors.NewBadRequestError(err.Error())
	}
	return f.fooService.Create(foo)
}

func (f *FooHandler) Update(c *fiber.Ctx) error {
	foo := new(model.Foo)
	if err := c.BodyParser(foo); err != nil {
		return errors.NewBadRequestError(err.Error())
	}
	return f.fooService.Update(foo)
}

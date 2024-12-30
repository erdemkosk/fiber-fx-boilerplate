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

// GetAll godoc
// @Summary Get all items
// @Description Get all foo items from the database
// @Tags foo
// @Accept json
// @Produce json
// @Success 200 {array} []model.Foo
// @Router /foo [get]
func (f *FooHandler) GetAll(c *fiber.Ctx) error {
	foos, err := f.fooService.GetAll()
	if err != nil {
		return errors.NewInternalError(err.Error())
	}
	return c.JSON(foos)
}

// GetById godoc
// @Summary Get a foo by id
// @Description Get a foo item by its ID
// @Tags foo
// @Accept json
// @Produce json
// @Param id path string true "Foo ID"
// @Success 200 {object} model.Foo
// @Failure 404 {object} error
// @Router /foo/{id} [get]
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

// Create godoc
// @Summary Create a new foo
// @Description Create a new foo item
// @Tags foo
// @Accept json
// @Produce json
// @Param foo body model.CreateFooRequest true "Foo object"
// @Success 201 {object} model.Foo
// @Failure 400 {object} error
// @Router /foo [post]
func (f *FooHandler) Create(c *fiber.Ctx) error {
	createRequest := new(model.CreateFooRequest)
	if err := c.BodyParser(createRequest); err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	foo := &model.Foo{
		Name:        createRequest.Name,
		Description: createRequest.Description,
	}

	if err := f.fooService.Create(foo); err != nil {
		return errors.NewInternalError(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
	})
}

// Update godoc
// @Summary Update a foo
// @Description Update an existing foo item by ID
// @Tags foo
// @Accept json
// @Produce json
// @Param id path string true "Foo ID"
// @Param foo body model.CreateFooRequest true "Foo object"
// @Success 200 {object} map[string]interface{} "{'success': true}"
// @Failure 400 {object} error "Bad request"
// @Failure 404 {object} error "Foo not found"
// @Failure 500 {object} error "Internal server error"
// @Router /foo/{id} [put]
func (f *FooHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return errors.NewBadRequestError("id is required")
	}

	updateRequest := new(model.CreateFooRequest)
	if err := c.BodyParser(updateRequest); err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	foo := &model.Foo{
		Name:        updateRequest.Name,
		Description: updateRequest.Description,
	}

	if err := f.fooService.Update(foo); err != nil {
		return errors.NewInternalError(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

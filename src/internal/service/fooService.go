package service

import (
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/model"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/repository"
)

type IFooService interface {
	GetAll() ([]*model.Foo, error)
	GetById(id string) (*model.Foo, error)
	Create(foo *model.Foo) error
	Update(foo *model.Foo) error
}

type FooService struct {
	fooRepository repository.IFooRepository
}

func NewFooService(fooRepository repository.IFooRepository) IFooService {
	return &FooService{
		fooRepository: fooRepository,
	}
}

func (f *FooService) GetAll() ([]*model.Foo, error) {
	return f.fooRepository.GetAll()
}

func (f *FooService) GetById(id string) (*model.Foo, error) {
	return f.fooRepository.GetById(id)
}

func (f *FooService) Create(foo *model.Foo) error {
	return f.fooRepository.Create(foo)
}

func (f *FooService) Update(foo *model.Foo) error {
	return f.fooRepository.Update(foo)
}

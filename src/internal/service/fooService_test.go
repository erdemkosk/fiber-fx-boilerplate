package service

import (
	"testing"

	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFooRepository struct {
	mock.Mock
}

func (m *MockFooRepository) GetAll() ([]*model.Foo, error) {
	args := m.Called()
	return args.Get(0).([]*model.Foo), args.Error(1)
}

func (m *MockFooRepository) GetById(id string) (*model.Foo, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Foo), args.Error(1)
}

func (m *MockFooRepository) Create(foo *model.Foo) error {
	args := m.Called(foo)
	return args.Error(0)
}

func (m *MockFooRepository) Update(foo *model.Foo) error {
	args := m.Called(foo)
	return args.Error(0)
}

func TestFooService_GetAll(t *testing.T) {
	mockRepo := new(MockFooRepository)
	service := NewFooService(mockRepo)

	expectedFoos := []*model.Foo{
		{ID: "1", Name: "Test Foo 1"},
		{ID: "2", Name: "Test Foo 2"},
	}

	mockRepo.On("GetAll").Return(expectedFoos, nil)

	foos, err := service.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedFoos, foos)
	mockRepo.AssertExpectations(t)
}

func TestFooService_GetById(t *testing.T) {
	mockRepo := new(MockFooRepository)
	service := NewFooService(mockRepo)

	expectedFoo := &model.Foo{ID: "1", Name: "Test Foo"}

	mockRepo.On("GetById", "1").Return(expectedFoo, nil)

	foo, err := service.GetById("1")

	assert.NoError(t, err)
	assert.Equal(t, expectedFoo, foo)
	mockRepo.AssertExpectations(t)
}

func TestFooService_Create(t *testing.T) {
	mockRepo := new(MockFooRepository)
	service := NewFooService(mockRepo)

	newFoo := &model.Foo{ID: "1", Name: "New Foo"}

	mockRepo.On("Create", newFoo).Return(nil)

	err := service.Create(newFoo)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestFooService_Update(t *testing.T) {
	mockRepo := new(MockFooRepository)
	service := NewFooService(mockRepo)

	updateFoo := &model.Foo{ID: "1", Name: "Updated Foo"}

	mockRepo.On("Update", updateFoo).Return(nil)

	err := service.Update(updateFoo)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

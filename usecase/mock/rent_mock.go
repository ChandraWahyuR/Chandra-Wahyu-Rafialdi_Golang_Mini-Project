package mock

import (
	"prototype/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockRentRepository struct {
	mock.Mock
}

func (m *MockRentRepository) PostRent(conf *domain.Rent) error {
	args := m.Called(conf)
	return args.Error(0)
}

func (m *MockRentRepository) GetAll() ([]*domain.Rent, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Rent), args.Error(1)
}

func (m *MockRentRepository) DeleteRent(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRentRepository) GetById(id int) (*domain.Rent, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Rent), args.Error(1)
}

func (m *MockRentRepository) UpdateRent(id int, rent *domain.Rent) (*domain.Rent, error) {
	args := m.Called(id, rent)
	return args.Get(0).(*domain.Rent), args.Error(1)
}

func (m *MockRentRepository) GetUserID(userID uuid.UUID) ([]*domain.Rent, error) {
	args := m.Called(userID)
	return args.Get(0).([]*domain.Rent), args.Error(1)
}

func (m *MockRentRepository) GetUnconfirmedRents(userID uuid.UUID) ([]*domain.Rent, error) {
	args := m.Called(userID)
	return args.Get(0).([]*domain.Rent), args.Error(1)
}

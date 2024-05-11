package mock

import (
	"prototype/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockRentConfirmRepository struct {
	mock.Mock
}

func (m *MockRentConfirmRepository) PostRentConfirm(conf *domain.RentConfirm) error {
	args := m.Called(conf)
	return args.Error(0)
}

func (m *MockRentConfirmRepository) GetAll() ([]*domain.RentConfirm, error) {
	args := m.Called()
	return args.Get(0).([]*domain.RentConfirm), args.Error(1)
}

func (m *MockRentConfirmRepository) GetById(id int) (*domain.RentConfirm, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.RentConfirm), args.Error(1)
}

func (m *MockRentConfirmRepository) ConfirmAdmin(id int, conf *domain.RentConfirm) (*domain.RentConfirm, error) {
	args := m.Called(id, conf)
	return args.Get(0).(*domain.RentConfirm), args.Error(1)
}

func (m *MockRentConfirmRepository) DeleteRentConfirm(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRentConfirmRepository) FindRentConfirmByUserId(userId uuid.UUID) ([]*domain.RentConfirm, error) {
	args := m.Called(userId)
	return args.Get(0).([]*domain.RentConfirm), args.Error(1)
}

func (m *MockRentConfirmRepository) CancelRentConfirmByUserId(ID int, userId uuid.UUID) error {
	args := m.Called(ID, userId)
	return args.Error(0)
}

func (m *MockRentConfirmRepository) GetAllInfoRental() ([]*domain.RentConfirm, error) {
	args := m.Called()
	return args.Get(0).([]*domain.RentConfirm), args.Error(1)
}

func (m *MockRentConfirmRepository) ConfirmReturnRental(ID int, conf *domain.RentConfirm) (*domain.RentConfirm, error) {
	args := m.Called(ID, conf)
	return args.Get(0).(*domain.RentConfirm), args.Error(1)
}

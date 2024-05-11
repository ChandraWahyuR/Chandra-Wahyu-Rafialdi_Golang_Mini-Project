package mock

import (
	"prototype/domain"

	"github.com/stretchr/testify/mock"
)

type MockEquipmentRepository struct {
	mock.Mock
}

func (m *MockEquipmentRepository) PostEquipment(equip *domain.Equipment) error {
	args := m.Called(equip)
	return args.Error(0)
}

func (m *MockEquipmentRepository) GetAll() ([]*domain.Equipment, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Equipment), args.Error(1)
}

func (m *MockEquipmentRepository) DeleteEquipment(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockEquipmentRepository) GetById(id int) (*domain.Equipment, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Equipment), args.Error(1)
}

func (m *MockEquipmentRepository) UpdateEquipment(id int, equipment *domain.Equipment) (*domain.Equipment, error) {
	args := m.Called(id, equipment)
	return args.Get(0).(*domain.Equipment), args.Error(1)
}

func (m *MockEquipmentRepository) UpdateQuantity(equipment *domain.Equipment) (*domain.Equipment, error) {
	args := m.Called(equipment)
	return args.Get(0).(*domain.Equipment), args.Error(1)
}

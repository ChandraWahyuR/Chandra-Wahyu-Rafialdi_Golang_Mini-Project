package usecase_test

import (
	"errors"
	"prototype/constant"
	"prototype/domain"
	"prototype/usecase"
	mock_data "prototype/usecase/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostEquipment(t *testing.T) {
	mockRepo := new(mock_data.MockEquipmentRepository)
	equipUseCase := usecase.NewEquipmentUseCase(mockRepo)

	equip := &domain.Equipment{Name: "Test Equipment", Price: 100, Stock: 10}
	mockRepo.On("PostEquipment", equip).Return(nil)
	result, err := equipUseCase.PostEquipment(equip)
	assert.NoError(t, err)
	assert.Equal(t, *equip, result)

	emptyEquip := &domain.Equipment{}
	mockRepo.On("PostEquipment", emptyEquip).Return(constant.ErrEmptyInput)
	_, err = equipUseCase.PostEquipment(emptyEquip)
	assert.Error(t, err, "Expected an error for empty input")
}

func TestGetAllequipment(t *testing.T) {
	mockRepo := new(mock_data.MockEquipmentRepository)
	equipUseCase := usecase.NewEquipmentUseCase(mockRepo)

	mockData := []*domain.Equipment{{ID: 1, Name: "Test 1", CategoryId: 1, Category: *&domain.CategoryEquipment{ID: 1, Name: "Test 1"}, Description: "Test 1", Image: "Test 1", Price: 100, Stock: 10}, {ID: 2, Name: "Test 1", CategoryId: 1, Category: *&domain.CategoryEquipment{ID: 1, Name: "Test 1"}, Description: "Test 1", Image: "Test 1", Price: 100, Stock: 10}}
	mockRepo.On("GetAll").Return(mockData, nil)
	result, err := equipUseCase.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, mockData, result)
}

func TestDeleteEquipment(t *testing.T) {
	mockRepo := new(mock_data.MockEquipmentRepository)
	equipUseCase := usecase.NewEquipmentUseCase(mockRepo)

	mockRepo.On("DeleteEquipment", 1).Return(nil)
	err := equipUseCase.DeleteEquipment(1)
	assert.NoError(t, err)
}

func TestGetById(t *testing.T) {
	mockRepo := new(mock_data.MockEquipmentRepository)
	equipUseCase := usecase.NewEquipmentUseCase(mockRepo)

	mockData := &domain.Equipment{ID: 1, Name: "Test Equipment", Price: 100, Stock: 10}
	mockRepo.On("GetById", 1).Return(mockData, nil)
	result, err := equipUseCase.GetById(1)
	assert.NoError(t, err)
	assert.Equal(t, mockData, result)
}

func TestUpdateEquipment(t *testing.T) {
	mockRepo := new(mock_data.MockEquipmentRepository)
	equipUseCase := usecase.NewEquipmentUseCase(mockRepo)

	mockData := &domain.Equipment{ID: 1, Name: "Test Equipment", Price: 100, Stock: 10}
	mockRepo.On("UpdateEquipment", 1, mock.Anything).Return(mockData, nil)
	result, err := equipUseCase.UpdateEquipment(1, &domain.Equipment{})
	assert.NoError(t, err)
	assert.Equal(t, mockData, result)
}

func TestUpdateQuantity(t *testing.T) {
	mockRepo := new(mock_data.MockEquipmentRepository)
	equipUseCase := usecase.NewEquipmentUseCase(mockRepo)

	mockData := &domain.Equipment{ID: 1, Name: "Test Equipment", Price: 100, Stock: 10}
	mockRepo.On("UpdateQuantity", mock.Anything).Return(mockData, nil)
	result, err := equipUseCase.UpdateQuantity(&domain.Equipment{})
	assert.NoError(t, err)
	assert.Equal(t, mockData, result)
}

func TestPostEquipment_Error(t *testing.T) {
	mockRepo := new(mock_data.MockEquipmentRepository)
	equipmentUseCase := usecase.NewEquipmentUseCase(mockRepo)

	equipment := &domain.Equipment{Name: "Test equipment", Price: 100, Stock: 10}
	errorMsg := "error during add equipment"

	mockRepo.On("PostEquipment", equipment).Return(errors.New(errorMsg))
	_, err := equipmentUseCase.PostEquipment(equipment)

	assert.EqualError(t, err, errorMsg)
}

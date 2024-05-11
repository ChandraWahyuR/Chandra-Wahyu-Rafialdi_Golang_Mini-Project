package usecase_test

import (
	"errors"
	"prototype/constant"
	"prototype/domain"
	"prototype/usecase"
	mock_data "prototype/usecase/mock"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPostRent(t *testing.T) {
	mockRepo := new(mock_data.MockRentRepository)
	equipUseCase := usecase.NewRentUseCase(mockRepo)

	equip := &domain.Rent{EquipmentId: 1, Quantity: 1}
	mockRepo.On("PostRent", equip).Return(nil)
	result, err := equipUseCase.PostRent(equip)
	assert.NoError(t, err)
	assert.Equal(t, *equip, result)

	emptyEquip := &domain.Rent{}
	mockRepo.On("PostRent", emptyEquip).Return(constant.ErrEmptyInput)
	_, err = equipUseCase.PostRent(emptyEquip)
	assert.Error(t, err, "Expected an error for empty input")
}

func TestGetAllRent(t *testing.T) {
	mockRepo := new(mock_data.MockRentRepository)
	equipUseCase := usecase.NewRentUseCase(mockRepo)

	mockData := []*domain.Rent{{ID: 1, Total: 100, Quantity: 1, UserId: uuid.UUID{}}, {ID: 2, Total: 100, Quantity: 2, UserId: uuid.UUID{}}}
	mockRepo.On("GetAll").Return(mockData, nil)
	result, err := equipUseCase.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, mockData, result)
}

func TestDeleteRent(t *testing.T) {
	mockRepo := new(mock_data.MockRentRepository)
	equipUseCase := usecase.NewRentUseCase(mockRepo)

	mockRepo.On("DeleteRent", 1).Return(nil)
	err := equipUseCase.DeleteRent(1)
	assert.NoError(t, err)
}

func TestGetByIdRent(t *testing.T) {
	mockRepo := new(mock_data.MockRentRepository)
	equipUseCase := usecase.NewRentUseCase(mockRepo)

	mockData := &domain.Rent{ID: 1, Total: 100, Quantity: 1, UserId: uuid.UUID{}}
	mockRepo.On("GetById", 1).Return(mockData, nil)
	result, err := equipUseCase.GetById(1)
	assert.NoError(t, err)
	assert.Equal(t, mockData, result)
}

func TestUpdateRent(t *testing.T) {
	mockRepo := new(mock_data.MockRentRepository)
	rentUseCase := usecase.NewRentUseCase(mockRepo)

	rent := &domain.Rent{ID: 1, Quantity: 5, Total: 250}
	existingRent := &domain.Rent{ID: 1, Quantity: 3, Total: 150}
	mockRepo.On("GetById", 1).Return(existingRent, nil)
	mockRepo.On("UpdateRent", 1, existingRent).Return(rent, nil)

	updatedRent, err := rentUseCase.UpdateRent(1, rent)
	assert.NoError(t, err)
	assert.Equal(t, rent, updatedRent)

	rent.Quantity = 0
	expectedErr := constant.ErrEmptyInput
	_, err = rentUseCase.UpdateRent(1, rent)
	assert.EqualError(t, err, expectedErr.Error())

	rent.Quantity = 5
	existingRent.RentConfirmID = 1
	expectedErr = constant.ErrUpdateData
	_, err = rentUseCase.UpdateRent(1, rent)
	assert.EqualError(t, err, expectedErr.Error())

	rent.Quantity = 5
	existingRent.RentConfirmID = 0
	mockRepo.On("UpdateRent", 1, existingRent).Return(nil, errors.New("some error"))
	_, err = rentUseCase.UpdateRent(1, rent)
	assert.Error(t, err)
}

func TestGetUserID(t *testing.T) {
	mockRepo := new(mock_data.MockRentRepository)
	rentUseCase := usecase.NewRentUseCase(mockRepo)
	userID := uuid.New()

	// Test valid case
	mockData := []*domain.Rent{{ID: 1}, {ID: 2}}
	mockRepo.On("GetUserID", userID).Return(mockData, nil)
	result, err := rentUseCase.GetUserID(userID)
	assert.NoError(t, err)
	assert.Equal(t, mockData, result)

	// Test repository error
	errMsg := "error getting rents by user ID"
	mockRepo.On("GetUserID", userID).Return(nil, errors.New(errMsg))
	_, err = rentUseCase.GetUserID(userID)
	assert.EqualError(t, err, errMsg)

	// Test empty result
	mockRepo.On("GetUserID", userID).Return([]*domain.Rent{}, nil)
	_, err = rentUseCase.GetUserID(userID)
	assert.EqualError(t, err, constant.ErrGetDataFromId.Error())
}

func TestGetUnconfirmedRents(t *testing.T) {
	mockRepo := new(mock_data.MockRentRepository)
	rentUseCase := usecase.NewRentUseCase(mockRepo)
	userID := uuid.New()

	// Test valid case
	mockData := []*domain.Rent{{ID: 1}, {ID: 2}}
	mockRepo.On("GetUnconfirmedRents", userID).Return(mockData, nil)
	result, err := rentUseCase.GetUnconfirmedRents(userID)
	assert.NoError(t, err)
	assert.Equal(t, mockData, result)

	// Test repository error
	errMsg := "error getting unconfirmed rents by user ID"
	mockRepo.On("GetUnconfirmedRents", userID).Return(nil, errors.New(errMsg))
	_, err = rentUseCase.GetUnconfirmedRents(userID)
	assert.EqualError(t, err, errMsg)
}

func TestPostRent_Error(t *testing.T) {
	mockRepo := new(mock_data.MockRentRepository)
	equipUseCase := usecase.NewRentUseCase(mockRepo)

	rent := &domain.Rent{EquipmentId: 1, Quantity: 1}
	errorMsg := "error posting rent"
	mockRepo.On("PostRent", rent).Return(errors.New(errorMsg))

	_, err := equipUseCase.PostRent(rent)
	assert.EqualError(t, err, errorMsg)
}

func TestUpdateRent_Error(t *testing.T) {
	mockRepo := new(mock_data.MockRentRepository)
	rentUseCase := usecase.NewRentUseCase(mockRepo)

	rent := &domain.Rent{ID: 1, Quantity: 5, Total: 250}
	existingRent := &domain.Rent{ID: 1, Quantity: 3, Total: 150, RentConfirmID: 1}

	mockRepo.On("GetById", 1).Return(existingRent, nil)

	_, err := rentUseCase.UpdateRent(1, rent)
	assert.EqualError(t, err, constant.ErrUpdateData.Error())
}

func TestGetAll_Error(t *testing.T) {
	mockRepo := new(mock_data.MockRentRepository)
	equipUseCase := usecase.NewRentUseCase(mockRepo)

	errorMsg := "error getting all rents"
	mockRepo.On("GetAll").Return([]*domain.Rent{}, errors.New(errorMsg))

	_, err := equipUseCase.GetAll()
	assert.EqualError(t, err, errorMsg)
}

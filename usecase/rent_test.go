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

func TestUpdateRent_Success(t *testing.T) {
	mockRepo := &mock_data.MockRentRepository{}
	rentUseCase := usecase.NewRentUseCase(mockRepo)

	id := 123
	rent := &domain.Rent{
		Quantity: 3,
		Total:    300,
	}

	existingRent := &domain.Rent{
		ID:            id,
		Quantity:      2,
		Total:         200,
		RentConfirmID: 0,
	}

	mockRepo.On("GetById", id).Return(existingRent, nil)
	mockRepo.On("UpdateRent", id, existingRent).Return(existingRent, nil)

	updatedRent, err := rentUseCase.UpdateRent(id, rent)

	assert.NoError(t, err)
	assert.NotNil(t, updatedRent)
	assert.Equal(t, rent.Quantity, updatedRent.Quantity)
	assert.Equal(t, rent.Total, updatedRent.Total)
}

func TestGetUserID_Success(t *testing.T) {
	mockRepo := &mock_data.MockRentRepository{}
	rentUseCase := usecase.NewRentUseCase(mockRepo)

	userID := uuid.New()

	expectedRents := []*domain.Rent{
		{ID: 1, UserId: userID},
		{ID: 2, UserId: userID},
	}

	mockRepo.On("GetUserID", userID).Return(expectedRents, nil)

	rents, err := rentUseCase.GetUserID(userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedRents, rents)
}

func TestGetUnconfirmedRents_Success(t *testing.T) {
	mockRepo := &mock_data.MockRentRepository{}
	rentUseCase := usecase.NewRentUseCase(mockRepo)

	userID := uuid.New()

	expectedRents := []*domain.Rent{
		{ID: 1, UserId: userID},
		{ID: 2, UserId: userID},
	}

	mockRepo.On("GetUnconfirmedRents", userID).Return(expectedRents, nil)

	rents, err := rentUseCase.GetUnconfirmedRents(userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedRents, rents)
}

func TestGetAll_Error(t *testing.T) {
	mockRepo := new(mock_data.MockRentRepository)
	equipUseCase := usecase.NewRentUseCase(mockRepo)

	errorMsg := "error getting all rents"
	mockRepo.On("GetAll").Return([]*domain.Rent{}, errors.New(errorMsg)).Once()

	_, err := equipUseCase.GetAll()
	assert.EqualError(t, err, errorMsg)
}

func TestUpdateRent_EmptyInput(t *testing.T) {
	mockRepo := &mock_data.MockRentRepository{}
	rentUseCase := usecase.NewRentUseCase(mockRepo)

	id := 123
	rent := &domain.Rent{}

	updatedRent, err := rentUseCase.UpdateRent(id, rent)

	assert.Error(t, err)
	assert.Nil(t, updatedRent)
	assert.Equal(t, constant.ErrEmptyInput, err)
}

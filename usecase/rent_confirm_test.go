package usecase_test

import (
	"prototype/constant"
	"prototype/domain"
	"prototype/usecase"
	"prototype/usecase/mock"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPostRentConfirm_Success(t *testing.T) {
	mockRepo := &mock.MockRentConfirmRepository{}
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)
	value := true
	conf := &domain.RentConfirm{
		PaymentMethod: "Credit Card",
		Delivery:      &value,
		Address:       "123 Main St",
		Rents: []domain.Rent{
			{Total: 100},
			{Total: 200},
		},
		Duration: 2,
	}

	mockRepo.On("PostRentConfirm", conf).Return(nil)

	result, err := rentConfirmUseCase.PostRentConfirm(conf)

	assert.NoError(t, err)
	assert.Equal(t, *conf, result)
}

func TestPostRentConfirm_EmptyInput(t *testing.T) {
	mockRepo := &mock.MockRentConfirmRepository{}
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)

	conf := &domain.RentConfirm{}

	result, err := rentConfirmUseCase.PostRentConfirm(conf)

	assert.Equal(t, domain.RentConfirm{}, result)
	assert.Equal(t, constant.ErrEmptyInput, err)
}

func TestPostRentConfirm_EmptyAddress(t *testing.T) {
	mockRepo := &mock.MockRentConfirmRepository{}
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)
	value := true
	conf := &domain.RentConfirm{
		PaymentMethod: "Credit Card",
		Delivery:      &value,
		Rents: []domain.Rent{
			{Total: 100},
			{Total: 200},
		},
		Duration: 2,
	}

	result, err := rentConfirmUseCase.PostRentConfirm(conf)

	assert.Equal(t, domain.RentConfirm{}, result)
	assert.Equal(t, constant.ErrEmptyAddress, err)
}

func TestGetAll_Success(t *testing.T) {
	mockRepo := &mock.MockRentConfirmRepository{}
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)
	value := true
	expectedRents := []*domain.RentConfirm{
		{ID: 1, PaymentMethod: "Credit Card", Delivery: &value},
		{ID: 2, PaymentMethod: "Cash", Delivery: &value},
	}

	mockRepo.On("GetAll").Return(expectedRents, nil)

	rents, err := rentConfirmUseCase.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedRents, rents)
}

func TestConfirmTestGetById(t *testing.T) {
	mockRepo := new(mock.MockRentConfirmRepository)
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)

	mockData := &domain.RentConfirm{ID: 1}
	mockRepo.On("GetById", 1).Return(mockData, nil)
	result, err := rentConfirmUseCase.GetById(1)
	assert.NoError(t, err)
	assert.Equal(t, mockData, result)

	mockRepo.AssertExpectations(t)
}

func TestDeleteRentConfirm_Success(t *testing.T) {
	mockRepo := &mock.MockRentConfirmRepository{}
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)

	id := 123

	mockRepo.On("DeleteRentConfirm", id).Return(nil)

	err := rentConfirmUseCase.DeleteRentConfirm(id)

	assert.NoError(t, err)
}

func TestFindRentConfirmByUserId_Success(t *testing.T) {
	mockRepo := &mock.MockRentConfirmRepository{}
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)

	userID := uuid.New()

	expectedRentConfirms := []*domain.RentConfirm{
		{ID: 1, UserId: userID},
		{ID: 2, UserId: userID},
	}

	mockRepo.On("FindRentConfirmByUserId", userID).Return(expectedRentConfirms, nil)

	rentConfirms, err := rentConfirmUseCase.FindRentConfirmByUserId(userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedRentConfirms, rentConfirms)
}

func TestCancelRentConfirmByUserId_Success(t *testing.T) {
	mockRepo := &mock.MockRentConfirmRepository{}
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)

	ID := 123
	userID := uuid.New()

	mockRepo.On("CancelRentConfirmByUserId", ID, userID).Return(nil)

	err := rentConfirmUseCase.CancelRentConfirmByUserId(ID, userID)

	assert.NoError(t, err)
}
func TestGetAllInfoRental_Success(t *testing.T) {
	mockRepo := &mock.MockRentConfirmRepository{}
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)

	value := true
	expectedRents := []*domain.RentConfirm{
		{ID: 1, PaymentMethod: "Credit Card", Delivery: &value},
		{ID: 2, PaymentMethod: "Cash", Delivery: &value},
	}

	mockRepo.On("GetAllInfoRental").Return(expectedRents, nil)

	rents, err := rentConfirmUseCase.GetAllInfoRental()

	assert.NoError(t, err)
	assert.Equal(t, expectedRents, rents)
}
func TestConfirmReturnRental_Success(t *testing.T) {
	mockRepo := &mock.MockRentConfirmRepository{}
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)

	ID := 123
	conf := &domain.RentConfirm{}

	mockRepo.On("ConfirmReturnRental", ID, conf).Return(conf, nil)

	confirmedRent, err := rentConfirmUseCase.ConfirmReturnRental(ID, conf)

	assert.NoError(t, err)
	assert.Equal(t, conf, confirmedRent)
}

func TestConfirmAdmin_Success(t *testing.T) {
	mockRepo := &mock.MockRentConfirmRepository{}
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)
	adminID := uuid.New()
	id := 123
	conf := &domain.RentConfirm{
		Status:  "Confirmed",
		AdminId: adminID,
	}

	existingRent := &domain.RentConfirm{
		ID:      id,
		Status:  "Pending",
		AdminId: adminID,
	}

	mockRepo.On("GetById", id).Return(existingRent, nil)
	mockRepo.On("ConfirmAdmin", id, existingRent).Return(existingRent, nil)

	updatedRent, err := rentConfirmUseCase.ConfirmAdmin(id, conf)

	assert.NoError(t, err)
	assert.NotNil(t, updatedRent)
	assert.Equal(t, conf.Status, updatedRent.Status)
	assert.Equal(t, conf.AdminId, updatedRent.AdminId)
}

package usecase_test

import (
	"errors"
	"prototype/constant"
	"prototype/domain"
	"prototype/usecase"
	"prototype/usecase/mock"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	mck "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestPostRentConfirm(t *testing.T) {
	mockRepo := new(mock.MockRentConfirmRepository)
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)

	conf := &domain.RentConfirm{
		PaymentMethod: "cash",
		Delivery:      new(bool),
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

	emptyConf := &domain.RentConfirm{}
	mockRepo.On("PostRentConfirm", emptyConf).Return(constant.ErrEmptyInput)
	_, err = rentConfirmUseCase.PostRentConfirm(emptyConf)
	assert.Error(t, err, "Expected an error for empty input")

	errMsg := "error posting rent confirm"
	mockRepo.On("PostRentConfirm", conf).Return(errors.New(errMsg))
	_, err = rentConfirmUseCase.PostRentConfirm(conf)
	assert.EqualError(t, err, constant.ErrInsertDatabase.Error())
}

func TestGetAll(t *testing.T) {
	mockRepo := new(mock.MockRentConfirmRepository)
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)

	// Test valid cxase
	mockData := []*domain.RentConfirm{{ID: 1}, {ID: 2}}
	mockRepo.On("GetAll").Return(mockData, nil)
	result, err := rentConfirmUseCase.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, mockData, result)

	errMsg := "error getting all rent confirms"
	mockRepo.On("GetAll").Return(nil, errors.New(errMsg))
	_, err = rentConfirmUseCase.GetAll()
	assert.EqualError(t, err, constant.ErrGetDatabase.Error())
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

func TestDeleteRentConfirm(t *testing.T) {
	mockRepo := new(mock.MockRentConfirmRepository)
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)

	mockRepo.On("DeleteRentConfirm", 1).Return(nil)
	err := rentConfirmUseCase.DeleteRentConfirm(1)
	assert.NoError(t, err)

	errMsg := "error deleting rent confirm"
	mockRepo.On("DeleteRentConfirm", 1).Return(errors.New(errMsg))
	err = rentConfirmUseCase.DeleteRentConfirm(1)
	assert.EqualError(t, err, constant.ErrDeleteData.Error())
}

func TestFindRentConfirmByUserId(t *testing.T) {
	mockRepo := new(mock.MockRentConfirmRepository)
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)
	userID := uuid.New()

	mockData := []*domain.RentConfirm{{ID: 1}, {ID: 2}}
	mockRepo.On("FindRentConfirmByUserId", userID).Return(mockData, nil)
	result, err := rentConfirmUseCase.FindRentConfirmByUserId(userID)
	assert.NoError(t, err)
	assert.Equal(t, mockData, result)

	errMsg := "error finding rent confirms by user ID"
	mockRepo.On("FindRentConfirmByUserId", userID).Return(nil, errors.New(errMsg))
	_, err = rentConfirmUseCase.FindRentConfirmByUserId(userID)
	assert.EqualError(t, err, errMsg)
}

func TestCancelRentConfirmByUserId(t *testing.T) {
	mockRepo := new(mock.MockRentConfirmRepository)
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)
	userID := uuid.New()

	mockRepo.On("CancelRentConfirmByUserId", 1, userID).Return(nil)
	err := rentConfirmUseCase.CancelRentConfirmByUserId(1, userID)
	assert.NoError(t, err)

	errMsg := "error canceling rent confirm by user ID"
	mockRepo.On("CancelRentConfirmByUserId", 1, userID).Return(errors.New(errMsg))
	err = rentConfirmUseCase.CancelRentConfirmByUserId(1, userID)
	assert.EqualError(t, err, errMsg)
}

func TestGetAllInfoRental(t *testing.T) {
	mockRepo := new(mock.MockRentConfirmRepository)
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)

	mockData := []*domain.RentConfirm{{ID: 1}, {ID: 2}}
	mockRepo.On("GetAllInfoRental").Return(mockData, nil)
	result, err := rentConfirmUseCase.GetAllInfoRental()
	assert.NoError(t, err)
	assert.Equal(t, mockData, result)

	errMsg := "error getting all rental info"
	mockRepo.On("GetAllInfoRental").Return(nil, errors.New(errMsg))
	_, err = rentConfirmUseCase.GetAllInfoRental()
	assert.EqualError(t, err, constant.ErrGetDatabase.Error())
}

func TestConfirmReturnRental(t *testing.T) {
	mockRepo := new(mock.MockRentConfirmRepository)
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)

	conf := &domain.RentConfirm{ID: 1, Status: "returned"}
	mockRepo.On("ConfirmReturnRental", 1, conf).Return(conf, nil)
	result, err := rentConfirmUseCase.ConfirmReturnRental(1, conf)
	assert.NoError(t, err)
	assert.Equal(t, conf, result)

	errMsg := "error confirming return of rental"
	mockRepo.On("ConfirmReturnRental", 1, conf).Return(nil, errors.New(errMsg))
	_, err = rentConfirmUseCase.ConfirmReturnRental(1, conf)
	assert.EqualError(t, err, errMsg)
}

func TestConfirmAdmin(t *testing.T) {
	mockRepo := new(mock.MockRentConfirmRepository)
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)

	mockRepo.On("GetById", mck.Anything).Return(&domain.RentConfirm{}, nil)

	conf := &domain.RentConfirm{ID: 1, Status: "confirmed", AdminId: uuid.New()}
	mockRepo.On("ConfirmAdmin", 1, mck.AnythingOfType("*domain.RentConfirm")).Return(conf, nil)
	result, err := rentConfirmUseCase.ConfirmAdmin(1, conf)
	assert.NoError(t, err)
	assert.Equal(t, conf, result)

	emptyStatusConf := &domain.RentConfirm{}
	_, err = rentConfirmUseCase.ConfirmAdmin(1, emptyStatusConf)
	assert.EqualError(t, err, constant.ErrEmptyStatus.Error())

	errMsg := "error confirming rent by admin"
	mockRepo.On("ConfirmAdmin", 1, mck.AnythingOfType("*domain.RentConfirm")).Return(nil, errors.New(errMsg))
	_, err = rentConfirmUseCase.ConfirmAdmin(1, conf)
	assert.EqualError(t, err, errMsg)
}

// Error
func TestPostRentConfirm_Error(t *testing.T) {
	mockRepo := new(mock.MockRentConfirmRepository)
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(mockRepo)

	rent := &domain.RentConfirm{
		PaymentMethod: "cash",
		Delivery:      new(bool),
		Rents: []domain.Rent{
			{Total: 100},
			{Total: 200},
		},
		Duration: 2,
	}
	errorMsg := "error posting rent confirm"
	mockRepo.On("PostRentConfirm", rent).Return(errors.New(errorMsg))

	_, err := rentConfirmUseCase.PostRentConfirm(rent)
	require.EqualError(t, err, errorMsg)
}

func TestGetAllRent_Error(t *testing.T) {
	mockRepo := new(mock.MockRentConfirmRepository)
	equipUseCase := usecase.NewRentConfirmUseCase(mockRepo)

	errorMsg := "error getting all rents"
	mockRepo.On("GetAll").Return([]*domain.RentConfirm{}, errors.New(errorMsg))

	_, err := equipUseCase.GetAll()
	assert.EqualError(t, err, errorMsg)
}

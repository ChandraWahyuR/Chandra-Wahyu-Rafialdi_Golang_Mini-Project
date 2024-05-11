package usecase_test

import (
	"errors"
	"prototype/constant"
	"prototype/domain"
	"prototype/usecase/mock"

	"prototype/usecase"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	mockRepo := new(mock.MockUserRepository)
	userUseCase := usecase.NewUserUseCase(mockRepo)

	user := &domain.User{Name: "Test User", Email: "test@example.com", Password: "password"}
	mockRepo.On("Register", user).Return(nil)
	result, err := userUseCase.Register(user)
	assert.NoError(t, err)
	assert.Equal(t, *user, result)

	emptyUser := &domain.User{}
	mockRepo.On("Register", emptyUser).Return(constant.ErrEmptyInput)
	_, err = userUseCase.Register(emptyUser)
	assert.Error(t, err, "Expected an error for empty input")

	errMsg := "error registering user"
	mockRepo.On("Register", user).Return(errors.New(errMsg))
	_, err = userUseCase.Register(user)
	assert.EqualError(t, err, constant.ErrInsertDatabase.Error())
}

func TestLogin(t *testing.T) {
	mockRepo := new(mock.MockUserRepository)
	userUseCase := usecase.NewUserUseCase(mockRepo)

	user := &domain.User{Email: "test@example.com", Password: "password"}
	mockRepo.On("Login", user).Return(nil)
	result, err := userUseCase.Login(user)
	assert.NoError(t, err)
	assert.Equal(t, *user, result)

	emptyUser := &domain.User{}
	mockRepo.On("Login", emptyUser).Return(constant.ErrEmptyInput)
	_, err = userUseCase.Login(emptyUser)
	assert.Error(t, err, "Expected an error for empty input")

	errMsg := "error logging in"
	mockRepo.On("Login", user).Return(errors.New(errMsg))
	_, err = userUseCase.Login(user)
	assert.EqualError(t, err, constant.ErrLogin.Error())
}

func TestGetByID(t *testing.T) {
	mockRepo := new(mock.MockUserRepository)
	userUseCase := usecase.NewUserUseCase(mockRepo)

	userID := uuid.New()
	mockUser := &domain.User{ID: userID, Name: "Test User", Email: "test@example.com"}

	mockRepo.On("GetByID", userID).Return(mockUser, nil)
	result, err := userUseCase.GetByID(userID)
	assert.NoError(t, err)
	assert.Equal(t, mockUser, result)

	errMsg := "error getting user by ID"
	mockRepo.On("GetByID", userID).Return(nil, errors.New(errMsg))
	_, err = userUseCase.GetByID(userID)
	assert.EqualError(t, err, errMsg)
}

func TestLogin_Error(t *testing.T) {
	mockRepo := new(mock.MockUserRepository)
	userUseCase := usecase.NewUserUseCase(mockRepo)

	user := &domain.User{Email: "test@example.com", Password: "password"}
	errorMsg := "error during login"

	mockRepo.On("Login", user).Return(errors.New(errorMsg))
	_, err := userUseCase.Login(user)

	assert.EqualError(t, err, errorMsg)
}
func TestRegister_Error(t *testing.T) {
	mockRepo := new(mock.MockUserRepository)
	userUseCase := usecase.NewUserUseCase(mockRepo)

	user := &domain.User{Name: "Test User", Email: "test@example.com", Password: "password"}
	errorMsg := "error during Register"

	mockRepo.On("Register", user).Return(errors.New(errorMsg))
	_, err := userUseCase.Register(user)

	assert.EqualError(t, err, errorMsg)
}

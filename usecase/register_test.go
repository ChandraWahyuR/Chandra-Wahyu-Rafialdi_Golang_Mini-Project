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

func TestRegister_Success(t *testing.T) {
	mockRepo := &mock.MockUserRepository{}
	userUseCase := usecase.NewUserUseCase(mockRepo)

	user := &domain.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	mockRepo.On("Register", user).Return(nil)

	createdUser, err := userUseCase.Register(user)

	assert.NoError(t, err)
	assert.Equal(t, user.Name, createdUser.Name)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.Password, createdUser.Password)
	mockRepo.AssertExpectations(t)
}

func TestRegister_EmptyInput(t *testing.T) {
	mockRepo := &mock.MockUserRepository{}
	userUseCase := usecase.NewUserUseCase(mockRepo)

	user := &domain.User{}

	createdUser, err := userUseCase.Register(user)

	assert.Equal(t, domain.User{}, createdUser)
	assert.Equal(t, constant.ErrEmptyInput, err)
}

func TestRegister_InsertError(t *testing.T) {
	mockRepo := &mock.MockUserRepository{}
	userUseCase := usecase.NewUserUseCase(mockRepo)

	user := &domain.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	mockRepo.On("Register", user).Return(errors.New("database error"))

	createdUser, err := userUseCase.Register(user)

	assert.Equal(t, domain.User{}, createdUser)
	assert.Equal(t, constant.ErrInsertDatabase, err)
	mockRepo.AssertExpectations(t)
}

func TestLogin_Success(t *testing.T) {
	mockRepo := &mock.MockUserRepository{}
	userUseCase := usecase.NewUserUseCase(mockRepo)

	user := &domain.User{Email: "test@example.com", Password: "password"}
	mockRepo.On("Login", user).Return(nil)

	result, err := userUseCase.Login(user)

	assert.NoError(t, err)
	assert.Equal(t, *user, result)
}

func TestLogin_EmptyInput(t *testing.T) {
	mockRepo := &mock.MockUserRepository{}
	userUseCase := usecase.NewUserUseCase(mockRepo)

	emptyUser := &domain.User{}
	_, err := userUseCase.Login(emptyUser)

	assert.Error(t, err)
	assert.Equal(t, constant.ErrEmptyInput, err)
}

func TestLogin_LoginError(t *testing.T) {
	mockRepo := &mock.MockUserRepository{}
	userUseCase := usecase.NewUserUseCase(mockRepo)

	user := &domain.User{Email: "test@example.com", Password: "password"}
	mockRepo.On("Login", user).Return(errors.New("error logging in"))

	_, err := userUseCase.Login(user)

	assert.EqualError(t, err, constant.ErrLogin.Error())
}
func TestGetByID_Success(t *testing.T) {
	mockRepo := &mock.MockUserRepository{}
	userUseCase := usecase.NewUserUseCase(mockRepo)

	userID := uuid.New()
	expectedUser := &domain.User{
		ID:    userID,
		Name:  "John Doe",
		Email: "john@example.com",
	}

	mockRepo.On("GetByID", userID).Return(expectedUser, nil)

	user, err := userUseCase.GetByID(userID)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser, user)
}

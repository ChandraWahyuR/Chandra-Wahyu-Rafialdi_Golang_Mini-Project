package mock

import (
	"prototype/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Register(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Login(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(userID uuid.UUID) (*domain.User, error) {
	args := m.Called(userID)
	return args.Get(0).(*domain.User), args.Error(1)
}

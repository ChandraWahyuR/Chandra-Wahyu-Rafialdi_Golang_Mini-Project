package usecase

import (
	"prototype/constant"
	"prototype/domain"

	"github.com/google/uuid"
)

type UserUseCase struct {
	repository domain.RepositoryInterface
}

func NewUserUseCase(repository domain.RepositoryInterface) *UserUseCase {
	return &UserUseCase{
		repository: repository,
	}
}

func (u *UserUseCase) Register(user *domain.User) (domain.User, error) {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return domain.User{}, constant.ErrEmptyInput
	}

	err := u.repository.Register(user)

	if err != nil {
		return domain.User{}, constant.ErrInsertDatabase
	}

	return *user, nil
}

func (u *UserUseCase) Login(user *domain.User) (domain.User, error) {
	if user.Email == "" || user.Password == "" {
		return domain.User{}, constant.ErrEmptyInput
	}
	err := u.repository.Login(user)

	if err != nil {
		return domain.User{}, constant.ErrLogin
	}

	return *user, nil
}

func (u *UserUseCase) GetByID(userID uuid.UUID) (*domain.User, error) {
	return u.repository.GetByID(userID)
}

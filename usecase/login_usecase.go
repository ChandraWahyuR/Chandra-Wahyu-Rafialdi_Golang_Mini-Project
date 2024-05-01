package usecase

import (
	"prototype/constant"
	"prototype/domain"
)

type LoginUseCase struct {
	repository domain.RepositoryInterface
}

func NewLoginUseCase(repository domain.RepositoryInterface) *LoginUseCase {
	return &LoginUseCase{
		repository: repository,
	}
}

func (u *LoginUseCase) Login(user *domain.User) (domain.User, error) {
	if user.Email == "" || user.Password == "" {
		return domain.User{}, constant.ErrEmptyInput
	}
	err := u.repository.Login(user)

	if err != nil {
		return domain.User{}, constant.ErrInsertDatabase
	}

	return *user, nil
}

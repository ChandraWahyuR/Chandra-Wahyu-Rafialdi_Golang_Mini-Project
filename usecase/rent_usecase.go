package usecase

import (
	"prototype/constant"
	"prototype/domain"

	"github.com/google/uuid"
)

type RentUseCase struct {
	repository domain.RentRepositoryInterface
}

func NewRentUseCase(repository domain.RentRepositoryInterface) *RentUseCase {
	return &RentUseCase{
		repository: repository,
	}
}

func (u *RentUseCase) PostRent(rent *domain.Rent) (domain.Rent, error) {
	if rent.Quantity == 0 {
		return domain.Rent{}, constant.ErrEmptyInput
	}
	err := u.repository.PostRent(rent)

	if err != nil {
		return domain.Rent{}, constant.ErrInsertDatabase
	}

	return *rent, nil
}

func (u *RentUseCase) GetAll() ([]*domain.Rent, error) {
	rent, err := u.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return rent, nil
}

func (u *RentUseCase) DeleteRent(id int) error {
	err := u.repository.DeleteRent(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *RentUseCase) GetById(id int) (*domain.Rent, error) {
	rent, err := u.repository.GetById(id)
	if err != nil {
		return nil, constant.ErrInsertDatabase
	}

	return rent, nil
}

func (u *RentUseCase) UpdateRent(id int, rent *domain.Rent) (*domain.Rent, error) {
	if rent.Quantity == 0 {
		return nil, constant.ErrEmptyInput
	}

	existingRent, err := u.repository.GetById(id)
	if err != nil {
		return nil, constant.ErrInsertDatabase
	}

	// if data is exist, data can be updated
	existingRent.Quantity = rent.Quantity
	existingRent.Total = rent.Total

	updatedRent, err := u.repository.UpdateRent(id, existingRent)
	if err != nil {
		return nil, err
	}
	return updatedRent, nil
}

func (u *RentUseCase) GetUserID(userID uuid.UUID) ([]*domain.Rent, error) {
	rents, err := u.repository.GetUserID(userID)
	if err != nil {
		return nil, err
	}
	return rents, nil
}

func (u *RentUseCase) GetUnconfirmedRents(userID uuid.UUID) ([]*domain.Rent, error) {
	rents, err := u.repository.GetUnconfirmedRents(userID)
	if err != nil {
		return nil, err
	}
	return rents, nil
}

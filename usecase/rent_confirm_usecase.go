package usecase

import (
	"prototype/constant"
	"prototype/domain"
)

type RentConfirmUseCase struct {
	repository domain.RentConfirmRepositoryInterface
}

func NewRentConfirmUseCase(repository domain.RentConfirmRepositoryInterface) *RentConfirmUseCase {
	return &RentConfirmUseCase{
		repository: repository,
	}
}

func (u *RentConfirmUseCase) PostRentConfirm(conf *domain.RentConfirm) (domain.RentConfirm, error) {
	if conf.PaymentMethod == "" || conf.Delivery == nil {
		return domain.RentConfirm{}, constant.ErrEmptyInput
	}

	if conf.Delivery != nil && *conf.Delivery == true {
		if conf.Address == "" {
			return domain.RentConfirm{}, constant.ErrEmptyAddress
		}
	}

	err := u.repository.PostRentConfirm(conf)
	if err != nil {
		return domain.RentConfirm{}, constant.ErrInsertDatabase
	}

	return *conf, nil
}

func (u *RentConfirmUseCase) GetAll() ([]*domain.RentConfirm, error) {
	rent, err := u.repository.GetAll()
	if err != nil {
		return nil, constant.ErrGetDatabase
	}

	return rent, nil
}

func (u *RentConfirmUseCase) GetById(id int) (*domain.RentConfirm, error) {
	rent, err := u.repository.GetById(id)
	if err != nil {
		return nil, constant.ErrFindData
	}

	return rent, nil
}

func (u *RentConfirmUseCase) ConfirmAdmin(id int, conf *domain.RentConfirm) (*domain.RentConfirm, error) {
	if conf.Status == "" {
		return nil, constant.ErrEmptyStatus
	}

	existingRent, err := u.repository.GetById(id)
	if err != nil {
		return nil, constant.ErrFindData
	}

	// if data is exist, data can be updated

	updatedRent, err := u.repository.ConfirmAdmin(id, existingRent)
	if err != nil {
		return nil, err
	}
	return updatedRent, nil
}

func (u *RentConfirmUseCase) DeleteRentConfirm(id int) error {
	err := u.repository.DeleteRentConfirm(id)
	if err != nil {
		return constant.ErrDeleteData
	}

	return nil
}

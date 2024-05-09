package usecase

import (
	"prototype/constant"
	"prototype/domain"

	"github.com/google/uuid"
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

	totalFee := 0
	for _, rent := range conf.Rents {
		totalFee += rent.Total
	}
	totalFee *= conf.Duration
	conf.Fee = totalFee

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

	existingRent.Status = conf.Status
	existingRent.AdminId = conf.AdminId

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

// New Feature
func (u *RentConfirmUseCase) FindRentConfirmByUserId(userId uuid.UUID) ([]*domain.RentConfirm, error) {
	rentConfirms, err := u.repository.FindRentConfirmByUserId(userId)
	if err != nil {
		return nil, constant.ErrById
	}
	if len(rentConfirms) == 0 {
		return nil, constant.ErrGetDataFromId
	}
	return rentConfirms, nil
}

func (u *RentConfirmUseCase) CancelRentConfirmByUserId(ID int, userId uuid.UUID) error {
	err := u.repository.CancelRentConfirmByUserId(ID, userId)
	if err != nil {
		return err
	}
	return nil
}

// Rental Info
func (u *RentConfirmUseCase) GetAllInfoRental() ([]*domain.RentConfirm, error) {
	rent, err := u.repository.GetAllInfoRental()
	if err != nil {
		return nil, constant.ErrGetDatabase
	}

	return rent, nil
}

func (u *RentConfirmUseCase) ConfirmReturnRental(ID int, conf *domain.RentConfirm) (*domain.RentConfirm, error) {
	confirmedRent, err := u.repository.ConfirmReturnRental(ID, conf)
	if err != nil {
		return nil, err
	}

	return confirmedRent, nil
}

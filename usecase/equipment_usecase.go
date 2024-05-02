package usecase

import (
	"prototype/constant"
	"prototype/domain"
)

type EquipmentUseCase struct {
	repository domain.EquipmentRepositoryInterface
}

func NewEquipmentUseCase(repository domain.EquipmentRepositoryInterface) *EquipmentUseCase {
	return &EquipmentUseCase{
		repository: repository,
	}
}

func (u *EquipmentUseCase) PostEquipment(equip *domain.Equipment) (domain.Equipment, error) {
	if equip.Name == "" || equip.Price == 0 {
		return domain.Equipment{}, constant.ErrEmptyInput
	}
	err := u.repository.PostEquipment(equip)

	if err != nil {
		return domain.Equipment{}, constant.ErrInsertDatabase
	}

	return *equip, nil
}

func (u *EquipmentUseCase) GetAll() ([]*domain.Equipment, error) {
	equipment, err := u.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return equipment, nil
}

func (u *EquipmentUseCase) DeleteEquipment(id int) error {
	err := u.repository.DeleteEquipment(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *EquipmentUseCase) GetById(id int) (*domain.Equipment, error) {
	equipment, err := u.repository.GetById(id)
	if err != nil {
		return nil, constant.ErrInsertDatabase
	}

	return equipment, nil
}

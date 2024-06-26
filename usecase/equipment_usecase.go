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
	if equip.Name == "" || equip.Price == 0 || equip.Stock == 0 {
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
		return nil, constant.ErrGetDatabase
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
		return nil, constant.ErrById
	}

	return equipment, nil
}

func (u *EquipmentUseCase) UpdateEquipment(id int, equipment *domain.Equipment) (*domain.Equipment, error) {

	updatedEquipment, err := u.repository.UpdateEquipment(id, equipment)
	if err != nil {
		return nil, err
	}

	return updatedEquipment, nil
}

func (uc *EquipmentUseCase) UpdateQuantity(equipment *domain.Equipment) (*domain.Equipment, error) {
	updatedEquipment, err := uc.repository.UpdateQuantity(equipment)
	if err != nil {
		return nil, err
	}
	return updatedEquipment, nil
}

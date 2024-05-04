package usecase

import (
	"prototype/constant"
	"prototype/domain"
)

type CategoryUseCase struct {
	repository domain.CategoryEquipmentRepositoryInterface
}

func NewCategoryUseCase(repository domain.CategoryEquipmentRepositoryInterface) *CategoryUseCase {
	return &CategoryUseCase{
		repository: repository,
	}
}

func (u *CategoryUseCase) PostCategoryEquipment(equip *domain.CategoryEquipment) (domain.CategoryEquipment, error) {
	if equip.Name == "" {
		return domain.CategoryEquipment{}, constant.ErrEmptyInput
	}
	err := u.repository.PostCategoryEquipment(equip)

	if err != nil {
		return domain.CategoryEquipment{}, constant.ErrInsertDatabase
	}

	return *equip, nil
}

func (u *CategoryUseCase) GetAll() ([]*domain.CategoryEquipment, error) {
	Categoryt, err := u.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return Categoryt, nil
}

func (u *CategoryUseCase) DeleteCategoryEquipment(id int) error {
	err := u.repository.DeleteCategoryEquipment(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *CategoryUseCase) GetById(id int) (*domain.CategoryEquipment, error) {
	rent, err := u.repository.GetById(id)
	if err != nil {
		return nil, constant.ErrInsertDatabase
	}

	return rent, nil
}

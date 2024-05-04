package drivers

import "prototype/domain"

type CategoryEquipment struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name string `json:"name"`
}

func FromCategoryEquipmentUseCase(eq *domain.CategoryEquipment) *CategoryEquipment {
	return &CategoryEquipment{
		ID:   eq.ID,
		Name: eq.Name,
	}
}

func (eq *CategoryEquipment) ToCategoryEquipmentUseCase() *domain.CategoryEquipment {
	return &domain.CategoryEquipment{
		ID:   eq.ID,
		Name: eq.Name,
	}
}

package response

import (
	"prototype/domain"
)

type EquipmentResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// CategoryId  int             `json:"category_id"`
	Category    CategoryDetails `json:"category"`
	Description string          `json:"description"`
	Image       string          `json:"image"`
	Price       int             `json:"price"`
	Stock       int             `json:"stock"`
}

type CategoryDetails struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func FromUseCase(equip *domain.Equipment) *EquipmentResponse {
	return &EquipmentResponse{
		ID:   equip.ID,
		Name: equip.Name,
		// CategoryId: equip.CategoryId,
		Category: CategoryDetails{
			Id:   equip.Category.ID,
			Name: equip.Category.Name,
		},
		Description: equip.Description,
		Image:       equip.Image,
		Price:       equip.Price,
		Stock:       equip.Stock,
	}
}

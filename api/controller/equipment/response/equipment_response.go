package response

import (
	"prototype/domain"
)

type EquipmentResponse struct {
	ID   int    `json:"id" example:"1"`
	Name string `json:"name" example:"Mesin Daur Ulang Kecil 1.5 Ton"`
	// CategoryId  int             `json:"category_id"`
	Category    CategoryDetails `json:"category" `
	Description string          `json:"description" example:"Mesin Daur Ulang Kecil 1.5 Ton"`
	Image       string          `json:"image" example:"https://cloudinary.com/photo/2016/03/31/15/32/robot-1295393_960_720.png"`
	Price       int             `json:"price" example:"1000000"`
	Stock       int             `json:"stock" example:"10"`
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

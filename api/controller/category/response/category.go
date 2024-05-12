package response

import "prototype/domain"

type CategoryResponse struct {
	ID   int    `json:"id" example:"1"`
	Name string `json:"name" example:"Daur Ulang"`
}

func FromUseCase(equip *domain.CategoryEquipment) *CategoryResponse {
	return &CategoryResponse{
		ID:   equip.ID,
		Name: equip.Name,
	}
}

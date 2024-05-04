package response

import "prototype/domain"

type CategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FromUseCase(equip *domain.CategoryEquipment) *CategoryResponse {
	return &CategoryResponse{
		ID:   equip.ID,
		Name: equip.Name,
	}
}

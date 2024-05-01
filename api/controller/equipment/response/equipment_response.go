package response

import "prototype/domain"

type EquipmentResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func FromUseCase(equip *domain.Equipment) *EquipmentResponse {
	return &EquipmentResponse{
		ID:          equip.ID,
		Name:        equip.Name,
		Category:    equip.Category,
		Description: equip.Description,
		Image:       equip.Image,
	}
}

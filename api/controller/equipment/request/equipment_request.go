package request

import "prototype/domain"

type EquipmentRequest struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
}

func (r *EquipmentRequest) ToEntities() *domain.Equipment {
	return &domain.Equipment{
		Name:        r.Name,
		Category:    r.Category,
		Description: r.Description,
		Image:       r.Image,
		Price:       r.Price,
	}
}

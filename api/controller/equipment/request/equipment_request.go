package request

import "prototype/domain"

type EquipmentRequest struct {
	Name        string `json:"name" form:"name"`
	CategoryId  int    `json:"category_id" form:"category_id"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
	Price       int    `json:"price" form:"price"`
	Stock       int    `json:"stock" form:"stock"`
}

func (r *EquipmentRequest) ToEntities() *domain.Equipment {
	return &domain.Equipment{
		Name:        r.Name,
		CategoryId:  r.CategoryId,
		Description: r.Description,
		Image:       r.Image,
		Price:       r.Price,
		Stock:       r.Stock,
	}
}

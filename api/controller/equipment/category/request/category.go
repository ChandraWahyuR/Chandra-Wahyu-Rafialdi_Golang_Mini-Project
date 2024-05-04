package request

import "prototype/domain"

type CategoryRequest struct {
	Name string `json:"name"`
}

func (r *CategoryRequest) ToEntities() *domain.CategoryEquipment {
	return &domain.CategoryEquipment{
		Name: r.Name,
	}
}

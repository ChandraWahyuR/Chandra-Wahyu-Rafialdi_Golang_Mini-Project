package request

import (
	"prototype/domain"
)

// Request http from user to prosesed by db
type RentRequest struct {
	EquipmentId int `json:"equipment_id"`
	Quantity    int `json:"quantity"`
	Total       int `json:"total"`
	Duration    int `json:"duration"`
}

func (r *RentRequest) ToEntities() *domain.Rent {
	return &domain.Rent{
		EquipmentId: r.EquipmentId,
		Quantity:    r.Quantity,
		Total:       r.Total,
		Duration:    r.Duration,
	}
}

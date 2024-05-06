package request

import "prototype/domain"

type RentConfirmRequest struct {
	PaymentMethod string `json:"payment_method"`
	Delivery      bool   `json:"delivery"`
	Address       string `json:"address"`
	Status        string `json:"status"`
	Duration      int    `json:"duration"`
	// Rents         []*EquipmentRequest `json:"rents"`
}

type EquipmentRequest struct {
	EquipmentID int `json:"equipment_id"`
	Total       int `json:"total"`
}

func (r *RentConfirmRequest) ToEntities() *domain.RentConfirm {
	return &domain.RentConfirm{
		PaymentMethod: r.PaymentMethod,
		Delivery:      &r.Delivery,
		Address:       r.Address,
		Status:        r.Status,
		Duration:      r.Duration,
		// Rents:         []*domain.Rent{},
	}
}

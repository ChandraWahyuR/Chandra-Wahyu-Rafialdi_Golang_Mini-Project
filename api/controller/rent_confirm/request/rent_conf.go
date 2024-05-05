package request

import "prototype/domain"

type RentConfirmRequest struct {
	PaymentMethod string `json:"payment_method"`
	Delivery      bool   `json:"delivery"`
	Duration      int    `json:"duration"`
	Address       string `json:"address"`
	Status        string `json:"status"`
}

func (r *RentConfirmRequest) ToEntities() *domain.RentConfirm {
	return &domain.RentConfirm{}
}

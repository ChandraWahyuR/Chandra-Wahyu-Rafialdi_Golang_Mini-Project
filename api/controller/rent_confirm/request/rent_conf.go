package request

import "prototype/domain"

type RentConfirmRequest struct {
	PaymentMethod string `json:"payment_method"`
	Delivery      bool   `json:"delivery"`
	Address       string `json:"address"`
	Status        string `json:"status"`
}

func (r *RentConfirmRequest) ToEntities() *domain.RentConfirm {
	return &domain.RentConfirm{
		PaymentMethod: r.PaymentMethod,
		Delivery:      &r.Delivery,
		Address:       r.Address,
		Status:        r.Status,
	}
}

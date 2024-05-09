package request

import "prototype/domain"

type PostConfirmRequest struct {
	PaymentMethod string                `json:"payment_method" form:"payment_method"`
	Delivery      bool                  `json:"delivery" form:"delivery"`
	Address       string                `json:"address" form:"address"`
	Status        string                `json:"status" form:"-"`
	Duration      int                   `json:"duration" form:"duration"`
	Rents         []*RentConfirmRequest `json:"rents" form:"-"`
}

func (r *PostConfirmRequest) ToEntities() *domain.RentConfirm {
	return &domain.RentConfirm{
		PaymentMethod: r.PaymentMethod,
		Delivery:      &r.Delivery,
		Address:       r.Address,
		Status:        r.Status,
		Duration:      r.Duration,
	}
}

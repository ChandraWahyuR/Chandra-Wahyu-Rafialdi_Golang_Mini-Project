package request

import (
	"prototype/domain"
	"time"
)

// Request http from user to prosesed by db
type RentRequest struct {
	Quantity  int
	Total     int
	DateStart time.Time
	Duration  int
}

func (r *RentRequest) ToEntities() *domain.Rent {
	return &domain.Rent{
		Quantity:  r.Quantity,
		Total:     r.Total,
		DateStart: r.DateStart,
		Duration:  r.Duration,
	}
}

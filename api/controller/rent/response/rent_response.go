package response

import (
	"prototype/domain"
	"time"
)

// Response from usecase to user

type RentResponse struct {
	ID        int       `json:"id"`
	Quantity  int       `json:"quantity"`
	Total     int       `json:"total"`
	DateStart time.Time `json:"date_start"`
	Duration  int       `json:"duration"`
}

func FromUseCase(rent *domain.Rent) *RentResponse {
	return &RentResponse{
		ID:        rent.ID,
		Quantity:  rent.Quantity,
		Total:     rent.Total,
		DateStart: rent.DateStart,
		Duration:  rent.Duration,
	}
}

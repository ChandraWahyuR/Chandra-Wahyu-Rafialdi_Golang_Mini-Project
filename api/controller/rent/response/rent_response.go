package response

import (
	"prototype/domain"
	"time"

	"github.com/google/uuid"
)

// Response from usecase to user

type RentResponse struct {
	ID          int       `json:"id"`
	UserId      uuid.UUID `json:"user_id"`
	EquipmentId int       `json:"equipment_id"`
	Quantity    int       `json:"quantity"`
	Total       int       `json:"total"`
	DateStart   time.Time `json:"date_start"`
	Duration    int       `json:"duration"`
}

func FromUseCase(rent *domain.Rent) *RentResponse {
	return &RentResponse{
		ID:          rent.ID,
		UserId:      rent.UserId,
		EquipmentId: rent.EquipmentId,
		Quantity:    rent.Quantity,
		Total:       rent.Total,
		DateStart:   rent.DateStart,
		Duration:    rent.Duration,
	}
}

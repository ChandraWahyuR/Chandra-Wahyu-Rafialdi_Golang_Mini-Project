package drivers

import (
	"prototype/domain"
	"time"

	"github.com/google/uuid"
)

type Rent struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement:true"`
	UserId      uuid.UUID `json:"user_id"`
	EquipmentId int       `json:"equipment_id"`
	Quantity    int       `json:"quantity"`
	Total       int       `json:"total"`
	DateStart   time.Time `json:"date_start"`
	UpdatedAt   time.Time `json:"updated_at"`
	Duration    int       `json:"duration"`
}

func FromRentUseCase(rent *domain.Rent) *Rent {
	return &Rent{
		ID:          rent.ID,
		UserId:      rent.UserId,
		EquipmentId: rent.EquipmentId,
		Quantity:    rent.Quantity,
		Total:       rent.Total,
		DateStart:   rent.DateStart,
		UpdatedAt:   rent.UpdatedAt,
		Duration:    rent.Duration,
	}
}

// for retrive data from gorm
func (rent *Rent) ToRentUseCase() *domain.Rent {
	return &domain.Rent{
		ID:          rent.ID,
		UserId:      rent.UserId,
		EquipmentId: rent.EquipmentId,
		Quantity:    rent.Quantity,
		Total:       rent.Total,
		DateStart:   rent.DateStart,
		UpdatedAt:   rent.UpdatedAt,
		Duration:    rent.Duration,
	}
}

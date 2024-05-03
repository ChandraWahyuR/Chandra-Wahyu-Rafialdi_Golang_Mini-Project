package drivers

import (
	"prototype/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rent struct {
	ID          int            `json:"id" gorm:"primaryKey;autoIncrement:true"`
	UserId      uuid.UUID      `json:"user_id"`
	EquipmentId int            `json:"equipment_id"`
	Equipment   Equipment      `json:"equipment"`
	Quantity    int            `json:"quantity"`
	Total       int            `json:"total"`
	DateStart   time.Time      `json:"date_start"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Duration    int            `json:"duration"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"` //buat nanti di table rent_equipment soft delete jadi data yang sudah di acc apa ditolak langsung dihapus
}

func FromRentUseCase(rent *domain.Rent) *Rent {
	return &Rent{
		ID:          rent.ID,
		UserId:      rent.UserId,
		EquipmentId: rent.EquipmentId,
		Equipment: Equipment{
			ID:          rent.Equipment.ID,
			Name:        rent.Equipment.Name,
			Category:    rent.Equipment.Category,
			Description: rent.Equipment.Description,
			Price:       rent.Equipment.Price,
		},
		Quantity:  rent.Quantity,
		Total:     rent.Total,
		DateStart: rent.DateStart,
		UpdatedAt: rent.UpdatedAt,
		Duration:  rent.Duration,
	}
}

// for retrive data from gorm
func (rent *Rent) ToRentUseCase() *domain.Rent {
	return &domain.Rent{
		ID:          rent.ID,
		UserId:      rent.UserId,
		EquipmentId: rent.EquipmentId,
		Equipment: domain.Equipment{
			ID:          rent.Equipment.ID,
			Name:        rent.Equipment.Name,
			Category:    rent.Equipment.Category,
			Description: rent.Equipment.Description,
			Price:       rent.Equipment.Price,
		},
		Quantity:  rent.Quantity,
		Total:     rent.Total,
		DateStart: rent.DateStart,
		UpdatedAt: rent.UpdatedAt,
		Duration:  rent.Duration,
	}
}

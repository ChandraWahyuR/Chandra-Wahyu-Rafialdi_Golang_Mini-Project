package drivers

import (
	"prototype/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rent struct {
	ID     int
	UserId uuid.UUID
	// RentConfirmID int
	EquipmentId int
	Equipment   Equipment
	Quantity    int
	Total       int
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func FromRentUseCase(rent *domain.Rent) *Rent {
	return &Rent{
		ID:          rent.ID,
		UserId:      rent.UserId,
		EquipmentId: rent.EquipmentId,
		Equipment: Equipment{
			ID:          rent.Equipment.ID,
			Name:        rent.Equipment.Name,
			CategoryId:  rent.Equipment.CategoryId,
			Description: rent.Equipment.Description,
			Image:       rent.Equipment.Image,
			Price:       rent.Equipment.Price,
			CreatedAt:   rent.Equipment.CreatedAt,
			UpdatedAt:   rent.Equipment.UpdatedAt,
			DeletedAt:   rent.Equipment.DeletedAt,
		},
		Quantity:  rent.Quantity,
		Total:     rent.Total,
		UpdatedAt: rent.UpdatedAt,
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
			CategoryId:  rent.Equipment.CategoryId,
			Description: rent.Equipment.Description,
			Image:       rent.Equipment.Image,
			Price:       rent.Equipment.Price,
			CreatedAt:   rent.Equipment.CreatedAt,
			UpdatedAt:   rent.Equipment.UpdatedAt,
			DeletedAt:   rent.Equipment.DeletedAt,
		},
		Quantity:  rent.Quantity,
		Total:     rent.Total,
		UpdatedAt: rent.UpdatedAt,
	}
}

package drivers

import (
	"prototype/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rent struct {
	ID            int            `json:"id" gorm:"primaryKey;autoIncrement:true"`
	UserId        uuid.UUID      `json:"user_id"`
	RentConfirmID int            `gorm:"many2many:rent_confirm_rents;"`
	EquipmentId   int            `json:"equipment_id"`
	Equipment     Equipment      `json:"equipment"`
	Quantity      int            `json:"quantity"`
	Total         int            `json:"total"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
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
			Description: rent.Equipment.Description,
			Image:       rent.Equipment.Image,
			Price:       rent.Equipment.Price,
			CreatedAt:   rent.Equipment.CreatedAt,
			UpdatedAt:   rent.Equipment.UpdatedAt,
			DeletedAt:   rent.Equipment.DeletedAt,
		},
		Quantity: rent.Quantity,
		Total:    rent.Total,
		// UpdatedAt: rent.UpdatedAt,
	}
}

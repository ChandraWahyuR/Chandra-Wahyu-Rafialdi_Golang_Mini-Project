package response

import (
	"prototype/domain"

	"github.com/google/uuid"
)

// Response from usecase to user

type RentResponse struct {
	ID          int              `json:"id"`
	UserId      uuid.UUID        `json:"user_id"`
	EquipmentId int              `json:"equipment_id"`
	Equipment   EquipmentDetails `json:"equipment"`
	Quantity    int              `json:"quantity"`
	Total       int              `json:"total"`
}

type EquipmentDetails struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
}

func FromUseCase(rent *domain.Rent) *RentResponse {
	return &RentResponse{
		ID:          rent.ID,
		UserId:      rent.UserId,
		EquipmentId: rent.EquipmentId,
		Equipment: EquipmentDetails{
			Name:        rent.Equipment.Name,
			Description: rent.Equipment.Description,
			Price:       rent.Equipment.Price,
			Image:       rent.Equipment.Image,
		},
		Quantity: rent.Quantity,
		Total:    rent.Total,
	}
}

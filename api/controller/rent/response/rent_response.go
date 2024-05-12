package response

import (
	"prototype/domain"

	"github.com/google/uuid"
)

// Response from usecase to user

type RentResponse struct {
	ID int `json:"id" example:"1"`
	// UserId uuid.UUID `json:"user_id"`
	User UserData `json:"user_data"`
	// EquipmentId int              `json:"equipment_id"`
	Equipment EquipmentDetails `json:"equipment"`
	Quantity  int              `json:"quantity" example:"1"`
	Total     int              `json:"total" example:"100000"`
}

type UserData struct {
	ID    uuid.UUID `json:"user_id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type EquipmentDetails struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
}

func FromUseCase(rent *domain.Rent) *RentResponse {
	return &RentResponse{
		ID: rent.ID,
		// UserId: rent.UserId,
		User: UserData{
			ID:    rent.User.ID,
			Name:  rent.User.Name,
			Email: rent.User.Email,
		},
		Equipment: EquipmentDetails{
			ID:          rent.Equipment.ID,
			Name:        rent.Equipment.Name,
			Description: rent.Equipment.Description,
			Price:       rent.Equipment.Price,
			Image:       rent.Equipment.Image,
		},
		Quantity: rent.Quantity,
		Total:    rent.Total,
	}
}

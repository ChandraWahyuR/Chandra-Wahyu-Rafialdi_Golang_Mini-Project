package response

import (
	"prototype/domain"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID       uuid.UUID `json:"id" example:"uuid"`
	Username string    `json:"username" example:"johndoe"`
	Name     string    `json:"name" example:"John Doe"`
	Email    string    `json:"email" example:"john@gmail.com"`
}

func FromUseCase(user *domain.User) *UserResponse {
	return &UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}
}

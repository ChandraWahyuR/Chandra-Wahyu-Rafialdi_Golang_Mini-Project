package response

import (
	"prototype/domain"

	"github.com/google/uuid"
)

type LoginResponse struct {
	ID    uuid.UUID `json:"id" example:"uuid"`
	Email string    `json:"email" example:"john@gmail.com"`
	Token string    `json:"token" form:"token" example:"token"`
}

func LoginUseCase(user *domain.User) *LoginResponse {
	return &LoginResponse{
		ID:    user.ID,
		Email: user.Email,
		Token: user.Token,
	}
}

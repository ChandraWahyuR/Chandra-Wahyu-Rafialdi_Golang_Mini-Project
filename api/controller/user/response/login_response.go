package response

import (
	"prototype/domain"

	"github.com/google/uuid"
)

type LoginResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Token string    `json:"token" form:"token"`
}

func LoginUseCase(user *domain.User) *LoginResponse {
	return &LoginResponse{
		ID:    user.ID,
		Email: user.Email,
		Token: user.Token,
	}
}

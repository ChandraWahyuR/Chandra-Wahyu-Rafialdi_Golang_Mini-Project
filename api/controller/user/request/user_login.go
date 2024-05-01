package request

import "prototype/domain"

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"roles"`
}

func (r *UserLogin) ToEntities() *domain.User {
	return &domain.User{
		Email:    r.Email,
		Password: r.Password,
	}
}

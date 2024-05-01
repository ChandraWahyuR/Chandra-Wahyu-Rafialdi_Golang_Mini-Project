package request

import "prototype/domain"

type UserRegister struct {
	Username        string `json:"username"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (r *UserRegister) ToEntities() *domain.User {
	return &domain.User{
		Username: r.Username,
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}

package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	Username    string
	Name        string
	Email       string
	Address     string
	PhoneNumber int
	Image       string
	Password    string
	Token       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type RepositoryInterface interface {
	Register(user *User) error
	Login(user *User) error
}

type UseCaseInterface interface {
	Register(user *User) (User, error)
	Login(user *User) (User, error)
}

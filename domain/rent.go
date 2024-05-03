package domain

import (
	"time"

	"github.com/google/uuid"
)

type Rent struct {
	ID          int
	UserId      uuid.UUID
	EquipmentId int
	Quantity    int
	Total       int
	DateStart   time.Time
	UpdatedAt   time.Time
	Duration    int
}

type RentRepositoryInterface interface {
	PostRent(rent *Rent) error
	GetAll() ([]*Rent, error)
	GetById(ID int) (*Rent, error)
	UpdateRent(ID int, rent *Rent) (*Rent, error)
	DeleteRent(ID int) error
}

type RentUseCaseInterface interface {
	PostRent(*Rent) (Rent, error)
	GetAll() ([]*Rent, error)
	GetById(ID int) (*Rent, error)
	UpdateRent(ID int, rent *Rent) (*Rent, error)
	DeleteRent(ID int) error
}

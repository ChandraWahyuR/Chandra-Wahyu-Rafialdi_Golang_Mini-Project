package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rent struct {
	ID            int
	UserId        uuid.UUID
	RentConfirmID int
	EquipmentId   int
	Equipment     Equipment
	Quantity      int
	Total         int
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type RentRepositoryInterface interface {
	PostRent(rent *Rent) error
	GetAll() ([]*Rent, error)
	GetById(ID int) (*Rent, error)
	UpdateRent(ID int, rent *Rent) (*Rent, error)
	DeleteRent(ID int) error

	GetUserID(userID uuid.UUID) ([]*Rent, error)
}

type RentUseCaseInterface interface {
	PostRent(*Rent) (Rent, error)
	GetAll() ([]*Rent, error)
	GetById(ID int) (*Rent, error)
	UpdateRent(ID int, rent *Rent) (*Rent, error)
	DeleteRent(ID int) error

	GetUserID(userID uuid.UUID) ([]*Rent, error)
}

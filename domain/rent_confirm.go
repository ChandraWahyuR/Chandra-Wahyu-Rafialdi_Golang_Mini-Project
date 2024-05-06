package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	StatusPending = "Pending"
	StatusAccept  = "Accept"
	StatusReject  = "Reject"
)

// `gorm:"many2many:rent_confirm_rents;"`
type RentConfirm struct {
	ID            int
	UserId        uuid.UUID
	Rents         []Rent `gorm:"foreignKey:RentConfirmID;"`
	Fee           int
	PaymentMethod string
	Delivery      *bool
	Address       string
	AdminId       uuid.UUID
	Status        string
	Duration      int
	DateStart     time.Time
	ReturnTime    time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type RentConfirmRepositoryInterface interface {
	PostRentConfirm(conf *RentConfirm) error
	GetAll() ([]*RentConfirm, error)
	GetById(ID int) (*RentConfirm, error)
	ConfirmAdmin(ID int, conf *RentConfirm) (*RentConfirm, error)
	DeleteRentConfirm(ID int) error
}

type RentConfirmUseCaseInterface interface {
	PostRentConfirm(*RentConfirm) (RentConfirm, error)
	GetAll() ([]*RentConfirm, error)
	GetById(ID int) (*RentConfirm, error)
	ConfirmAdmin(ID int, rent *RentConfirm) (*RentConfirm, error)
	DeleteRentConfirm(ID int) error
}

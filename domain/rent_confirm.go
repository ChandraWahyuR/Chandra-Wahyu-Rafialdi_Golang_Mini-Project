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

type RentConfirm struct {
	ID            int
	UserId        uuid.UUID
	RentId        int
	Rent          Rent
	Fee           int
	PaymentMethod string
	Delivery      *bool
	Address       string
	AdminId       uuid.UUID //pas diconfirm, extrak jwt ambil admin id put ke admin id
	Status        string
	ReturnTime    time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type RentConfirmRepositoryInterface interface {
	PostRentConfirm(conf *RentConfirm) error                      //users
	GetAll() ([]*RentConfirm, error)                              //admin
	GetById(ID int) (*RentConfirm, error)                         //admin
	ConfirmAdmin(ID int, conf *RentConfirm) (*RentConfirm, error) //admin
	DeleteRentConfirm(ID int) error                               //admin
}

type RentConfirmUseCaseInterface interface {
	PostRent(*RentConfirm) (RentConfirm, error)                   //users
	GetAll() ([]*RentConfirm, error)                              //admin
	GetById(ID int) (*RentConfirm, error)                         //admin
	ConfirmAdmin(ID int, rent *RentConfirm) (*RentConfirm, error) //admin
	DeleteRent(ID int) error                                      //admin
}

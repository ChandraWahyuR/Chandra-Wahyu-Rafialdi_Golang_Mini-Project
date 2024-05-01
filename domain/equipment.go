package domain

import (
	"time"

	"gorm.io/gorm"
)

type Equipment struct {
	ID          int
	Name        string
	Category    string
	Description string
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

// Logic related to db
type EquipmentRepositoryInterface interface {
	PostEquipment(equip *Equipment) error
	GetAll() ([]*Equipment, error)
	GetById(ID int) (*Equipment, error)
	DeleteEquipment(ID int) error
}

// Logic related to user activity
type EquipmentUseCaseInterface interface {
	PostEquipment(*Equipment) (Equipment, error)
	GetAll() ([]*Equipment, error)
	GetById(ID int) (*Equipment, error)
	DeleteEquipment(ID int) error
}
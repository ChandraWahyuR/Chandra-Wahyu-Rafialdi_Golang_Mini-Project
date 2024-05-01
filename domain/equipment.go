package domain

import (
	"time"
)

type Equipment struct {
	ID          int
	Name        string
	Category    Categories
	Description string
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type Categories struct {
	ID   int
	Name string
}

type EquipmentRepositoryInterface interface {
}

type EquipmentUseCaseInterface interface {
}

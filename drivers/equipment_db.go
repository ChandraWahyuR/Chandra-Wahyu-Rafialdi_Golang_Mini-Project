package drivers

import (
	"prototype/domain"
	"time"

	"gorm.io/gorm"
)

type Equipment struct {
	ID          int               `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name        string            `json:"name"`
	CategoryId  int               `json:"category_id"`
	Category    CategoryEquipment `json:"category"`
	Description string            `json:"description"`
	Image       string            `json:"image"`
	Price       int               `json:"price"`
	Stock       int               `json:"stock"`
	CreatedAt   time.Time         `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time         `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt    `json:"deleted_at" gorm:"index"`
}

// Conversi objek for domain layer, later save to gorm
func FromEquipmentUseCase(eq *domain.Equipment) *Equipment {
	return &Equipment{
		ID:          eq.ID,
		Name:        eq.Name,
		CategoryId:  eq.CategoryId,
		Description: eq.Description,
		Image:       eq.Image,
		Price:       eq.Price,
		Stock:       eq.Stock,
		CreatedAt:   eq.CreatedAt,
		UpdatedAt:   eq.UpdatedAt,
	}
}

// for retrive data from gorm
func (eq *Equipment) ToEquipmentUseCase() *domain.Equipment {
	return &domain.Equipment{
		ID:          eq.ID,
		Name:        eq.Name,
		CategoryId:  eq.CategoryId,
		Description: eq.Description,
		Image:       eq.Image,
		Price:       eq.Price,
		Stock:       eq.Stock,
		CreatedAt:   eq.CreatedAt,
		UpdatedAt:   eq.UpdatedAt,
		Category: domain.CategoryEquipment{
			ID:   eq.Category.ID,
			Name: eq.Category.Name,
		},
	}
}

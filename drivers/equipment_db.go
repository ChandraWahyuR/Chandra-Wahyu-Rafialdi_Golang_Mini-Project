package drivers

import (
	"prototype/domain"
	"time"

	"gorm.io/gorm"
)

type Equipment struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name        string `json:"name"`
	CategoryID  int    `json:"category_id"`
	Description string `json:"description" `
	Image       string
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Categories struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name string `json:"name"`
}

// Conversi objek for domain layer, later save to gorm
func FromEquipmentUseCase(eq *domain.Equipment) *Equipment {
	return &Equipment{
		ID:          eq.ID,
		Name:        eq.Name,
		CategoryID:  eq.Category.ID,
		Description: eq.Description,
		Image:       eq.Image,
		CreatedAt:   eq.CreatedAt,
		UpdatedAt:   eq.UpdatedAt,
	}
}

// for retrive data from gorm
func (eq *Equipment) ToEquipmentUseCase() *domain.Equipment {
	var category domain.Categories
	return &domain.Equipment{
		ID:          eq.ID,
		Name:        eq.Name,
		Category:    category,
		Description: eq.Description,
		Image:       eq.Image,
		CreatedAt:   eq.CreatedAt,
		UpdatedAt:   eq.UpdatedAt,
	}
}

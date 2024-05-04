package drivers

import (
	"prototype/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ID        uuid.UUID      `json:"id" gorm:"primaryKey;autoIncrement:true"`

type User struct {
	ID          uuid.UUID      `json:"id" gorm:"primaryKey"`
	Username    string         `json:"username"  gorm:"unique;not null"`
	Name        string         `json:"name"`
	Email       string         `json:"email" gorm:"unique;not null"`
	Address     string         `json:"address"`
	PhoneNumber int            `json:"phone_number"`
	Image       string         `json:"image"`
	Password    string         `json:"password"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func FromUseCase(user *domain.User) *User {
	return &User{
		ID:          uuid.New(),
		Username:    user.Username,
		Name:        user.Name,
		Email:       user.Email,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		Image:       user.Image,
		Password:    user.Password,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

func (user *User) ToUseCase() *domain.User {
	return &domain.User{
		ID:          user.ID,
		Username:    user.Username,
		Name:        user.Name,
		Email:       user.Email,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		Image:       user.Image,
		Password:    user.Password,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

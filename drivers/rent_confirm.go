package drivers

import (
	"prototype/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RentConfirm struct {
	ID            int            `json:"id" gorm:"primaryKey;autoIncrement:true"`
	UserId        uuid.UUID      `json:"user_id"`
	Fee           int            `json:"fee"`
	PaymentMethod string         `json:"payment_method"`
	Delivery      bool           `json:"delevery"`
	Address       string         `json:"address"`
	AdminId       uuid.UUID      `json:"admin_id"` //pas diconfirm, extrak jwt ambil admin id put ke admin id
	Status        string         `json:"status"`
	Duration      int            `json:"duration"`
	DateStart     time.Time      `json:"date_start"`
	ReturnTime    time.Time      `json:"return_time"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Rents         []*RentDetail  `gorm:"foreignKey:RentConfirmID"`
}

type RentDetail struct {
	ID            int
	UserId        uuid.UUID
	RentConfirmID int
	EquipmentId   int
	Quantity      int
	Total         int
}

func FromRentConfirmUseCase(conf *domain.RentConfirm) *RentConfirm {
	return &RentConfirm{
		ID:            conf.ID,
		UserId:        conf.UserId,
		Fee:           conf.Fee,
		PaymentMethod: conf.PaymentMethod,
		Address:       conf.Address,
		AdminId:       conf.AdminId,
		Status:        conf.Status,
		Duration:      conf.Duration,
		CreatedAt:     conf.CreatedAt,
		UpdatedAt:     conf.UpdatedAt,
	}
}

func (conf *RentConfirm) ToRentConfirmUseCase() *domain.RentConfirm {
	return &domain.RentConfirm{
		ID:            conf.ID,
		UserId:        conf.UserId,
		Fee:           conf.Fee,
		PaymentMethod: conf.PaymentMethod,
		Address:       conf.Address,
		AdminId:       conf.AdminId,
		Status:        conf.Status,
		Duration:      conf.Duration,
		ReturnTime:    conf.ReturnTime,
		CreatedAt:     conf.CreatedAt,
		UpdatedAt:     conf.UpdatedAt,
	}
}
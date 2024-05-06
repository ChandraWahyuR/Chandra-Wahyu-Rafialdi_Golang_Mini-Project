package response

import (
	"prototype/domain"
	"time"

	"github.com/google/uuid"
)

type RentConfirmRespond struct {
	ID            int           `json:"id"`
	UserId        uuid.UUID     `json:"user_id"`
	Rent          []RentDetails `json:"rent" gorm:"foreignKey:RentConfirmID"`
	Duration      int           `json:"duration"`
	Fee           int           `json:"fee"`
	PaymentMethod string        `json:"payment_method"`
	Delivery      bool          `json:"delivery"`
	Address       string        `json:"address"`
	AdminId       uuid.UUID     `json:"admin_id"`
	Status        string        `json:"status"`
	DateStart     time.Time     `json:"date_start"`
	ReturnTime    *time.Time    `json:"return_time"` // ini awal user kirim kosong, nanti pas admin confirm baru isi
}

type RentDetails struct {
	EquipmentId int `json:"equipment_id"`
	Total       int `json:"total"`
}

func FromUseCase(conf *domain.RentConfirm) *RentConfirmRespond {
	rentDetails := make([]RentDetails, len(conf.Rents))
	for i, rent := range conf.Rents {
		rentDetails[i] = RentDetails{
			EquipmentId: rent.EquipmentId,
			Total:       rent.Total,
		}
	}

	return &RentConfirmRespond{
		ID:            conf.ID,
		UserId:        conf.UserId,
		Rent:          rentDetails,
		Duration:      conf.Duration,
		Fee:           conf.Fee,
		PaymentMethod: conf.PaymentMethod,
		Delivery:      *conf.Delivery,
		Address:       conf.Address,
		AdminId:       conf.AdminId,
		Status:        conf.Status,
		DateStart:     conf.DateStart,
		ReturnTime:    &conf.ReturnTime,
	}
}

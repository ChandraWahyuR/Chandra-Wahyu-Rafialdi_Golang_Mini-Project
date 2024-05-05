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
	var returnTime *time.Time
	if !conf.ReturnTime.IsZero() {
		returnTime = &conf.ReturnTime
	}
	totalFee := 0
	rent := make([]RentDetails, len(conf.Rents))
	for i, r := range conf.Rents {
		rent[i] = RentDetails{
			EquipmentId: r.EquipmentId,
			Total:       r.Total,
		}
		totalFee += r.Total * conf.Duration
	}
	conf.Status = domain.StatusPending

	return &RentConfirmRespond{
		ID:            conf.ID,
		UserId:        conf.UserId,
		Rent:          rent,
		Duration:      conf.Duration,
		Fee:           totalFee,
		PaymentMethod: conf.PaymentMethod,
		Delivery:      *conf.Delivery,
		Address:       conf.Address,
		AdminId:       conf.AdminId,
		Status:        conf.Status,
		DateStart:     conf.DateStart,
		ReturnTime:    returnTime,
	}
}

package response

import (
	"prototype/domain"
	"time"

	"github.com/google/uuid"
)

type RentConfirmRequest struct {
	ID            int           `json:"id"`
	UserId        uuid.UUID     `json:"user_id"`
	Rent          []RentDetails `json:"rent"`
	Fee           int           `json:"fee"`
	PaymentMethod string        `json:"payment_method"`
	Delivery      bool          `json:"delivery"`
	Address       string        `json:"address"`
	AdminId       uuid.UUID     `json:"admin_id"`
	Status        string        `json:"status"`
	ReturnTime    *time.Time    `json:"return_time"` // ini awal user kirim kosong, nanti pas admin confirm baru isi
}

type RentDetails struct {
	EquipmentId int `json:"equipment_id"`
	Total       int `json:"total"`
}

func FromUseCase(conf *domain.RentConfirm) *RentConfirmRequest {
	var returnTime *time.Time
	if !conf.ReturnTime.IsZero() {
		returnTime = &conf.ReturnTime
	}

	// Hitung total biaya
	totalFee := 0
	rent := make([]RentDetails, len(conf.Rents))
	for i, r := range conf.Rents {
		rent[i] = RentDetails{
			EquipmentId: r.EquipmentId,
			Total:       r.Total,
		}
		// Tambahkan biaya sewa ke total biaya
		totalFee += r.Total
	}

	return &RentConfirmRequest{
		ID:            conf.ID,
		UserId:        conf.UserId,
		Rent:          rent,
		Fee:           totalFee, // Setel total biaya
		PaymentMethod: conf.PaymentMethod,
		Delivery:      *conf.Delivery,
		Address:       conf.Address,
		AdminId:       conf.AdminId,
		Status:        conf.Status,
		ReturnTime:    returnTime,
	}
}

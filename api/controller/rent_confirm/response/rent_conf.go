package response

import (
	"prototype/domain"
	"time"

	"github.com/google/uuid"
)

type RentConfirmRequest struct {
	ID            int         `json:"id"`
	UserId        uuid.UUID   `json:"user_id"`
	RentId        int         `json:"rent_id"`
	Rent          RentDetails `json:"rent"`
	Fee           int         `json:"fee"`
	PaymentMethod string      `json:"payment_method"`
	Delivery      bool        `json:"delivery"`
	Address       string      `json:"address"`
	AdminId       uuid.UUID   `json:"admin_id"`
	Status        string      `json:"status"`
	ReturnTime    time.Time   `json:"return_time"` // ini awal user kirim kosong, nanti pas admin confirm baru isi
}

type RentDetails struct {
	EquipmentId int `json:"equipment_id"`
	Total       int `json:"total"`
}

func FromUseCase(conf *domain.RentConfirm) *RentConfirmRequest {
	return &RentConfirmRequest{
		ID:     conf.ID,
		UserId: conf.UserId,
		RentId: conf.Rent.ID,
		Rent: RentDetails{
			EquipmentId: conf.Rent.EquipmentId,
			Total:       conf.Rent.Total,
		},
		Fee:           conf.Fee,
		PaymentMethod: conf.PaymentMethod,
		Delivery:      *conf.Delivery,
		Address:       conf.Address,
		AdminId:       conf.AdminId,
		Status:        conf.Status,
		ReturnTime:    conf.ReturnTime,
	}
}

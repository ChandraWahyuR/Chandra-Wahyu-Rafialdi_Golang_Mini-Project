package response

import (
	"prototype/domain"
	"time"

	"github.com/google/uuid"
)

type RentConfirmRespond struct {
	ID int `json:"id"`
	// UserId        uuid.UUID     `json:"user_id"`
	User          UserData      `json:"user_data"`
	Rent          []RentDetails `json:"rent" gorm:"foreignKey:RentConfirmID"`
	Duration      int           `json:"duration"`
	Fee           int           `json:"fee"`
	PaymentMethod string        `json:"payment_method"`
	Delivery      bool          `json:"delivery"`
	Address       string        `json:"address"`
	AdminId       uuid.UUID     `json:"admin_id"`
	Status        string        `json:"status"`
	DateStart     time.Time     `json:"date_start"`
	ReturnTime    time.Time     `json:"return_time"` // ini awal user kirim kosong, nanti pas admin confirm baru isi
}

type RentDetails struct {
	EquipmentId   int    `json:"equipment_id"`
	EquipmentName string `json:"name"`
	Total         int    `json:"total"`
}

type UserData struct {
	ID    uuid.UUID `json:"user_id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func FromUseCase(conf *domain.RentConfirm) *RentConfirmRespond {
	rentDetails := make([]RentDetails, len(conf.Rents))
	for i, rent := range conf.Rents {
		equipmentName := ""
		if rent.Equipment.ID != 0 {
			equipmentName = rent.Equipment.Name
		}

		rentDetails[i] = RentDetails{
			EquipmentId:   rent.EquipmentId,
			EquipmentName: equipmentName,
			Total:         rent.Total,
		}
	}

	var delivery bool
	if conf.Delivery != nil {
		delivery = *conf.Delivery
	} else {
		delivery = false
	}

	var timeStart time.Time
	if conf.Status != domain.StatusAccept {
		timeStart = conf.DateStart
	} else {
		timeStart = time.Now()
	}
	return &RentConfirmRespond{
		ID: conf.ID,
		// UserId: conf.UserId,
		Rent: rentDetails,
		User: UserData{
			ID:    conf.User.ID,
			Name:  conf.User.Name,
			Email: conf.User.Email,
		},
		Duration:      conf.Duration,
		Fee:           conf.Fee,
		PaymentMethod: conf.PaymentMethod,
		Delivery:      delivery,
		Address:       conf.Address,
		AdminId:       conf.AdminId,
		Status:        conf.Status,
		DateStart:     timeStart,
		ReturnTime:    conf.ReturnTime,
	}
}

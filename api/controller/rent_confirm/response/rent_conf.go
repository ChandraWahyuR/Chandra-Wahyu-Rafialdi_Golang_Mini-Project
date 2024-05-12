package response

import (
	"prototype/domain"
	"time"

	"github.com/google/uuid"
)

type RentConfirmRespond struct {
	ID            int           `json:"id" example:"1"`
	User          UserData      `json:"user_data"`
	Rent          []RentDetails `json:"rent" gorm:"foreignKey:RentConfirmID"`
	Duration      int           `json:"duration" example:"1"`
	Fee           int           `json:"fee" example:"100000"`
	PaymentMethod string        `json:"payment_method" example:"http://cloudinary.com/photo/2016/03/31/15/32/robot-1295393_960_720.png"`
	Delivery      bool          `json:"delivery" example:"true"`
	Address       string        `json:"address" example:"Jl. Setiabudi No. 1, Jakarta, Indonesia"`
	AdminId       uuid.UUID     `json:"admin_id" example:"uuid"`
	Status        string        `json:"status" example:"pending"`
	DateStart     time.Time     `json:"date_start" example:"2024-00-00 00:00:00"`
	ReturnTime    time.Time     `json:"return_time" example:"2024-00-00 00:00:00"`
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
		ID:   conf.ID,
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

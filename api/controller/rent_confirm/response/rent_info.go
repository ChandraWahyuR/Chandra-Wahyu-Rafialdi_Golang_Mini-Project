package response

import (
	"prototype/domain"
	"time"
)

type RentalInfoRespond struct {
	ID          int               `json:"id"`
	User        UserData          `json:"user_data"`
	Rent        []RentDetailsInfo `json:"rent"`
	Duration    int               `json:"duration"`
	Fee         int               `json:"fee"`
	Address     string            `json:"address"`
	ReturnTime  time.Time         `json:"return_time"`
	Description string            `json:"description"`
}

type RentDetailsInfo struct {
	EquipmentId   int    `json:"equipment_id"`
	EquipmentName string `json:"name"`
	Total         int    `json:"total"`
	Quantity      int    `json:"quantity"`
}

func FromUseCaseInfo(conf *domain.RentConfirm) *RentalInfoRespond {
	rentDetails := make([]RentDetailsInfo, len(conf.Rents))
	for i, rent := range conf.Rents {
		equipmentName := ""
		if rent.Equipment.ID != 0 {
			equipmentName = rent.Equipment.Name
		}

		rentDetails[i] = RentDetailsInfo{
			EquipmentId:   rent.EquipmentId,
			EquipmentName: equipmentName,
			Total:         rent.Total,
			Quantity:      rent.Quantity,
		}
	}

	return &RentalInfoRespond{
		ID:   conf.ID,
		Rent: rentDetails,
		User: UserData{
			ID:    conf.User.ID,
			Name:  conf.User.Name,
			Email: conf.User.Email,
		},
		Duration:    conf.Duration,
		Fee:         conf.Fee,
		Address:     conf.Address,
		ReturnTime:  conf.ReturnTime,
		Description: conf.Description,
	}
}

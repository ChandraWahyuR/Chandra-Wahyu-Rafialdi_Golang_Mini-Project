package repository

import (
	"prototype/domain"
	"prototype/drivers"
	"time"

	"gorm.io/gorm"
)

type RentConfirmRepo struct {
	DB *gorm.DB
}

func NewRentConfirmRepo(db *gorm.DB) *RentConfirmRepo {
	return &RentConfirmRepo{DB: db}
}

func (r *RentConfirmRepo) PostRentConfirm(conf *domain.RentConfirm) error {
	totalFee := 0
	for _, rent := range conf.Rents {
		totalFee += rent.Total
	}

	conf.Fee = totalFee
	conf.Status = domain.StatusPending
	resp := drivers.FromRentConfirmUseCase(conf)
	if err := r.DB.Create(&resp).Error; err != nil {
		return err
	}

	conf.ID = resp.ID
	return nil
}

func (r *RentConfirmRepo) GetAll() ([]*domain.RentConfirm, error) {
	var conf []*domain.RentConfirm
	if err := r.DB.Preload("Rent").Find(&conf).Error; err != nil {
		return nil, err
	}

	return conf, nil
}

func (r *RentConfirmRepo) GetById(id int) (*domain.RentConfirm, error) {
	db := &drivers.RentConfirm{}
	if err := r.DB.Preload("Equipment").First(db, id).Error; err != nil {
		return nil, err
	}
	return db.ToRentConfirmUseCase(), nil
}

func (r *RentConfirmRepo) ConfirmAdmin(id int, conf *domain.RentConfirm) (*domain.RentConfirm, error) {
	db := &drivers.Rent{}
	if err := r.DB.Where("id = ?", id).First(&db).Error; err != nil {
		return nil, err
	}

	resp := drivers.FromRentConfirmUseCase(conf)
	if err := r.DB.Save(resp).Error; err != nil {
		return nil, err
	}

	return resp.ToRentConfirmUseCase(), nil
}

func (r *RentConfirmRepo) DeleteRentConfirm(id int) error {
	db := &drivers.RentConfirm{}
	if err := r.DB.Where("id = ?", id).Delete(&db).Error; err != nil {
		return err
	}
	// Soft delete
	if err := r.DB.Model(db).Update("deleted_at", time.Now()).Where("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

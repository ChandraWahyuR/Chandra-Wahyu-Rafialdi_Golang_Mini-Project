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
	resp := drivers.FromRentConfirmUseCase(conf)
	if err := r.DB.Create(&resp).Error; err != nil {
		return err
	}

	conf.ID = resp.ID
	return nil
}

func (r *RentConfirmRepo) GetAll() ([]*domain.RentConfirm, error) {
	var conf []*domain.RentConfirm
	status := domain.StatusPending
	if err := r.DB.Where("status = ?", status).Find(&conf).Error; err != nil {
		return nil, err
	}

	for _, c := range conf {
		rents, err := r.getRentsByConfirmID(c.ID)
		if err != nil {
			return nil, err
		}
		c.Rents = rents
	}

	return conf, nil
}

func (r *RentConfirmRepo) getRentsByConfirmID(confirmID int) ([]*domain.Rent, error) {
	var rents []*domain.Rent
	if err := r.DB.Where("rent_confirm_id = ?", confirmID).Find(&rents).Error; err != nil {
		return nil, err
	}
	return rents, nil
}

func (r *RentConfirmRepo) GetById(id int) (*domain.RentConfirm, error) {
	db := &drivers.RentConfirm{}
	if err := r.DB.First(db, id).Error; err != nil {
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

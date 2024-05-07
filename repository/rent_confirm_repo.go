package repository

import (
	"prototype/domain"
	"prototype/drivers"
	"time"

	"github.com/google/uuid"
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
	// query := r.DB.Table("rents").Where("user_id = ?", resp.UserId).Update("delete_at", time.Now())
	// user_id dan deleted_at != null

	// Bagian confrim hapus, atau dideleted at
	conf.ID = resp.ID
	return nil
}

func (r *RentConfirmRepo) GetAll() ([]*domain.RentConfirm, error) {
	var rentConfirms []*domain.RentConfirm
	status := "Pending"
	if err := r.DB.Unscoped().Preload("Rents").Where("status = ?", status).Find(&rentConfirms).Error; err != nil {
		return nil, err
	}

	return rentConfirms, nil
}
func (r *RentConfirmRepo) GetById(id int) (*domain.RentConfirm, error) {
	db := &drivers.RentConfirm{}
	if err := r.DB.Unscoped().Preload("Rents").First(db, id).Error; err != nil {
		return nil, err
	}

	return db.ToRentConfirmUseCase(), nil
}

func (r *RentConfirmRepo) ConfirmAdmin(id int, conf *domain.RentConfirm) (*domain.RentConfirm, error) {
	db := &drivers.RentConfirm{}
	if err := r.DB.Unscoped().Where("id = ?", id).First(&db).Error; err != nil {
		return nil, err
	}

	db.Status = conf.Status
	db.DateStart = time.Now()
	if err := r.DB.Save(db).Error; err != nil {
		return nil, err
	}

	// Soft Delete
	if err := r.DB.Model(&drivers.Rent{}).Where("rent_confirm_id = ?", id).Update("deleted_at", time.Now()).Error; err != nil {
		return nil, err
	}

	return db.ToRentConfirmUseCase(), nil
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

// New Feature
// Get Confirmation about ren By User ID
func (r *RentConfirmRepo) FindRentConfirmByUserId(id uuid.UUID) ([]*domain.RentConfirm, error) {
	var db []*drivers.RentConfirm
	if err := r.DB.Unscoped().Preload("Rents").Where("user_id = ?", id).Find(&db).Error; err != nil {
		return nil, err
	}

	conf := make([]*domain.RentConfirm, len(db))
	for i, value := range db {
		conf[i] = value.ToRentConfirmUseCase()
	}
	return conf, nil
}

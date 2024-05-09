package repository

import (
	"errors"
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

	if conf.Delivery != nil {
		resp.Delivery = *conf.Delivery
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
	if err := r.DB.Unscoped().Preload("User").Preload("Rents", func(db *gorm.DB) *gorm.DB { return db.Preload("Equipment") }).Where("status = ?", status).Find(&rentConfirms).Error; err != nil {
		return nil, err
	}

	return rentConfirms, nil
}
func (r *RentConfirmRepo) GetById(id int) (*domain.RentConfirm, error) {
	db := &drivers.RentConfirm{}
	if err := r.DB.Unscoped().Preload("User").Preload("Rents", func(db *gorm.DB) *gorm.DB { return db.Preload("Equipment") }).First(db, id).Error; err != nil {
		return nil, err
	}

	return db.ToRentConfirmUseCase(), nil
}

func (r *RentConfirmRepo) ConfirmAdmin(id int, conf *domain.RentConfirm) (*domain.RentConfirm, error) {
	db := &drivers.RentConfirm{}
	if err := r.DB.Unscoped().Preload("User").Preload("Rents", func(db *gorm.DB) *gorm.DB { return db.Preload("Equipment") }).Where("id = ?", id).First(&db).Error; err != nil {
		return nil, err
	}
	now := time.Now()
	returnDate := now.Add(time.Duration(conf.Duration) * 7 * 24 * time.Hour)

	// Keterangan
	description := domain.StatusNotReturn
	// Save to db
	if conf.Status == domain.StatusAccept {
		db.Status = conf.Status
		db.AdminId = conf.AdminId
		db.DateStart = now
		db.ReturnTime = returnDate
		db.Description = description
	} else {
		db.Status = conf.Status
		db.AdminId = conf.AdminId
	}

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
	if err := r.DB.Unscoped().Preload("User").Preload("Rents", func(db *gorm.DB) *gorm.DB { return db.Preload("Equipment") }).Where("user_id = ?", id).Find(&db).Error; err != nil {
		return nil, err
	}

	conf := make([]*domain.RentConfirm, len(db))
	for i, value := range db {
		conf[i] = value.ToRentConfirmUseCase()
	}
	return conf, nil
}

// User Cancel RentConfirm
func (r *RentConfirmRepo) CancelRentConfirmByUserId(id int, userId uuid.UUID) error {
	var rentConfirm drivers.RentConfirm
	if err := r.DB.Where("id = ?", id).Find(&rentConfirm).Error; err != nil {
		return err
	}

	// Validate data to userid from jwt
	if rentConfirm.ID == 0 {
		return errors.New("rent_confirm not found")
	}
	if rentConfirm.UserId != userId {
		return errors.New("you are not authorized to cancel this rent_confirm")
	}

	if rentConfirm.Status != domain.StatusPending {
		return errors.New("rent_confirm cannot be cancelled because it is already confirmed")
	}

	// Update Data from rent to
	if err := r.DB.Model(&drivers.Rent{}).Where("rent_confirm_id = ?", id).Update("rent_confirm_id", 0).Error; err != nil {
		return err
	}

	// Hard delete
	if err := r.DB.Unscoped().Delete(&rentConfirm).Error; err != nil {
		return err
	}

	return nil
}

// Rental Info
func (r *RentConfirmRepo) GetAllInfoRental() ([]*domain.RentConfirm, error) {
	var rentConfirms []*domain.RentConfirm
	if err := r.DB.
		Unscoped().
		Preload("User").
		Preload("Rents", func(db *gorm.DB) *gorm.DB { return db.Preload("Equipment") }).
		Where("description <> '' AND status = ?", domain.StatusAccept).
		Find(&rentConfirms).
		Error; err != nil {
		return nil, err
	}
	return rentConfirms, nil
}

func (r *RentConfirmRepo) ConfirmReturnRental(id int, conf *domain.RentConfirm) (*domain.RentConfirm, error) {
	db := &drivers.RentConfirm{}
	if err := r.DB.Unscoped().Preload("User").Preload("Rents", func(db *gorm.DB) *gorm.DB { return db.Preload("Equipment") }).Where("id = ?", id).First(&db).Error; err != nil {
		return nil, err
	}

	db.Description = conf.Description
	if err := r.DB.Save(db).Error; err != nil {
		return nil, err
	}

	return db.ToRentConfirmUseCase(), nil
}

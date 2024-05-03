package repository

import (
	"prototype/domain"
	"prototype/drivers"
	"time"

	"gorm.io/gorm"
)

type RentRepo struct {
	DB *gorm.DB
}

func NewRentRepo(db *gorm.DB) *RentRepo {
	return &RentRepo{DB: db}
}

func (r *RentRepo) PostRent(rent *domain.Rent) error {
	resp := drivers.FromRentUseCase(rent)
	if err := r.DB.Create(&resp).Error; err != nil {
		return err
	}

	rent.ID = resp.ID
	return nil
}

func (r *RentRepo) GetAll() ([]*domain.Rent, error) {
	var rents []*domain.Rent
	if err := r.DB.Preload("Equipment").Find(&rents).Error; err != nil {
		return nil, err
	}

	return rents, nil
}

func (r *RentRepo) GetById(id int) (*domain.Rent, error) {
	db := &drivers.Rent{}
	if err := r.DB.Preload("Equipment").First(db, id).Error; err != nil {
		return nil, err
	}
	return db.ToRentUseCase(), nil
}

func (r *RentRepo) DeleteRent(id int) error {
	db := &drivers.Rent{}
	if err := r.DB.Where("id = ?", id).Delete(&db).Error; err != nil {
		return err
	}
	// Soft delete
	if err := r.DB.Model(db).Update("deleted_at", time.Now()).Where("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *RentRepo) UpdateRent(id int, rent *domain.Rent) (*domain.Rent, error) {
	db := &drivers.Rent{}
	if err := r.DB.Where("id = ?", id).First(&db).Error; err != nil {
		return nil, err
	}

	// Update
	resp := drivers.FromRentUseCase(rent)
	if err := r.DB.Save(resp).Error; err != nil {
		return nil, err
	}

	return resp.ToRentUseCase(), nil
}

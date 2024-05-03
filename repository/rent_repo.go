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
	var db []*drivers.Rent
	if err := r.DB.Find(&db).Error; err != nil {
		return nil, err
	}

	var rent []*domain.Rent
	for _, value := range db {
		rent = append(rent, value.ToRentUseCase())
	}

	return rent, nil
}

func (r *RentRepo) GetById(id int) (*domain.Rent, error) {
	db := &drivers.Rent{}
	if err := r.DB.First(db, id).Error; err != nil {
		return nil, err
	}
	return db.ToRentUseCase(), nil
}

func (r *RentRepo) DeleteRent(id int) error {
	db := &drivers.Rent{}
	if err := r.DB.Where("id = ?", id).Delete(&db).Error; err != nil {
		return err
	}

	if err := r.DB.Model(db).Update("deleted_at", time.Now()).Error; err != nil {
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

package repository

import (
	"prototype/domain"
	"prototype/drivers"
	"time"

	"gorm.io/gorm"
)

type EquipmentRepo struct {
	DB *gorm.DB
}

func NewEquipmentRepo(db *gorm.DB) *EquipmentRepo {
	return &EquipmentRepo{DB: db}
}

func (r *EquipmentRepo) PostEquipment(equip *domain.Equipment) error {
	resp := drivers.FromEquipmentUseCase(equip)
	if err := r.DB.Preload("Category").Create(&resp).Error; err != nil {
		return err
	}

	equip.ID = resp.ID
	return nil
}

func (r *EquipmentRepo) GetAll() ([]*domain.Equipment, error) {
	var db []*domain.Equipment
	if err := r.DB.Preload("Category").Find(&db).Error; err != nil {
		return nil, err
	}

	return db, nil
}

func (r *EquipmentRepo) GetById(id int) (*domain.Equipment, error) {
	db := &drivers.Equipment{}
	if err := r.DB.First(db, id).Error; err != nil {
		return nil, err
	}
	return db.ToEquipmentUseCase(), nil
}

func (r *EquipmentRepo) DeleteEquipment(id int) error {
	db := &drivers.Equipment{}
	if err := r.DB.Where("id = ?", id).Delete(&db).Error; err != nil {
		return err
	}

	if err := r.DB.Model(db).Update("deleted_at", time.Now()).Where("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

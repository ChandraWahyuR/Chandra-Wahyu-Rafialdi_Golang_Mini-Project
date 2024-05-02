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
	if err := r.DB.Create(&resp).Error; err != nil {
		return err
	}

	equip.ID = resp.ID
	return nil
}

func (r *EquipmentRepo) GetAll() ([]*domain.Equipment, error) {
	var db []*drivers.Equipment
	if err := r.DB.Find(&db).Error; err != nil {
		return nil, err
	}

	var equipment []*domain.Equipment
	for _, equip := range db {
		equipment = append(equipment, equip.ToEquipmentUseCase())
	}

	return equipment, nil
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

	if err := r.DB.Model(db).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}

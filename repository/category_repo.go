package repository

import (
	"prototype/domain"
	"prototype/drivers"
	"time"

	"gorm.io/gorm"
)

type CategoryRepo struct {
	DB *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{DB: db}
}

func (r *CategoryRepo) PostCategoryEquipment(category *domain.CategoryEquipment) error {
	resp := drivers.FromCategoryEquipmentUseCase(category)
	if err := r.DB.Create(&resp).Error; err != nil {
		return err
	}

	category.ID = resp.ID
	return nil
}

func (r *CategoryRepo) GetAll() ([]*domain.CategoryEquipment, error) {
	var db []*drivers.CategoryEquipment
	if err := r.DB.Find(&db).Error; err != nil {
		return nil, err
	}

	var Category []*domain.CategoryEquipment
	for _, equip := range db {
		Category = append(Category, equip.ToCategoryEquipmentUseCase())
	}

	return Category, nil
}

func (r *CategoryRepo) DeleteCategoryEquipment(id int) error {
	db := &drivers.CategoryEquipment{}
	if err := r.DB.Where("id = ?", id).Delete(&db).Error; err != nil {
		return err
	}

	if err := r.DB.Model(db).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepo) GetById(id int) (*domain.CategoryEquipment, error) {
	db := &drivers.CategoryEquipment{}
	if err := r.DB.First(db, id).Error; err != nil {
		return nil, err
	}
	return db.ToCategoryEquipmentUseCase(), nil
}

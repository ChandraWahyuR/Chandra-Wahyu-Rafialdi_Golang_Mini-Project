package repository

import (
	"prototype/constant"
	"prototype/domain"
	"prototype/drivers"
	"prototype/utils"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Register(user *domain.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	userDb := drivers.FromUseCase(user)

	if err := r.DB.Create(&userDb).Error; err != nil {
		return err
	}

	user.ID = userDb.ID
	return nil
}

func (r *UserRepo) Login(user *domain.User) error {
	userDb := &drivers.User{}
	if err := r.DB.Where("email = ?", user.Email).First(&userDb).Error; err != nil {
		return err
	}

	if !utils.CheckPasswordHash(user.Password, userDb.Password) {
		return constant.ErrAddUsersPassword
	}

	user.ID = userDb.ID
	return nil
}

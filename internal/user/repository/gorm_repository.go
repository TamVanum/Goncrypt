package repository

import (
	"github.com/tamVanum/goncrypt/config"
	"github.com/tamVanum/goncrypt/internal/user/model"
)

type userGormRepo struct{}

func NewUserGormRepo() UserRepository {
	return &userGormRepo{}
}

func (r *userGormRepo) FindAll() ([]model.User, error) {
	var users []model.User
	err := config.DB.Find(&users).Error
	return users, err
}

func (r *userGormRepo) Create(user *model.User) error {
	return config.DB.Create(user).Error
}

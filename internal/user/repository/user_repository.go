package repository

import "github.com/tamVanum/goncrypt/internal/user/model"

type UserRepository interface {
	Create(user *model.User) error
	FindAll() ([]model.User, error)
}

package service

import (
	"github.com/tamVanum/goncrypt/internal/user/model"
	"github.com/tamVanum/goncrypt/internal/user/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetUsers() ([]model.User, error) {
	return s.repo.FindAll()
}

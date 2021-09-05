package service

import (
	"access-management/models"
	"access-management/pkg/domain/user"
)

type userService struct {
	repo user.Repository
}

func User(repo user.Repository) user.Service {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Get() (models.User, error) {
	return s.repo.Get()
}

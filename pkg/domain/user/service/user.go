package service

import (
	"access-management/pkg/domain/user"
	"access-management/pkg/models"
)

type userService struct {
	userGateway user.Gateway
}

func User(userGateway user.Gateway) user.Service {
	return &userService{
		userGateway: userGateway,
	}
}

func (s *userService) Get() (models.User, error) {
	return s.userGateway.Get()
}

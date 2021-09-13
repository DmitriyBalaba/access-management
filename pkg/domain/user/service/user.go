package service

import (
	"access-management/pkg/domain/company"
	"access-management/pkg/domain/user"
	"access-management/pkg/models"
)

type userService struct {
	userGateway    user.Gateway
	companyGateway company.Gateway
}

func User(userGateway user.Gateway, companyGateway company.Gateway) user.Service {
	return &userService{
		userGateway:    userGateway,
		companyGateway: companyGateway,
	}
}

func (s *userService) Get() (models.User, error) {
	s.companyGateway.Get()
	return s.userGateway.Get()
}

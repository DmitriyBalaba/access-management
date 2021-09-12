package service

import (
	"access-management/pkg/domain/auth"
	"access-management/pkg/models"
)

type authService struct {
	userGateway auth.Gateway
	validator   auth.Validator
}

func Auth(userGateway auth.Gateway) auth.Service {
	return &authService{
		userGateway: userGateway,
		validator:   &authValidator{},
	}
}

func (s *authService) Signup(u *models.User) (*models.User, error) {

	if err := s.validator.ValidateUser(u); err != nil {
		return nil, err
	}

	if err := s.validator.ValidateCompany(u.Company); err != nil {
		return nil, err
	}

	return s.userGateway.Signup(u)
}

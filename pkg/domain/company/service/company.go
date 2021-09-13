package service

import (
	"access-management/pkg/domain/company"
	"access-management/pkg/domain/user"
	"access-management/pkg/models"
)

type companyService struct {
	companyGateway company.Gateway
	userGateway    user.Gateway
}

func Company(companyGateway company.Gateway, userGateway user.Gateway) company.Service {
	return &companyService{
		userGateway:    userGateway,
		companyGateway: companyGateway,
	}
}

func (s *companyService) Get() (models.Company, error) {
	s.userGateway.Get()
	return s.companyGateway.Get()
}

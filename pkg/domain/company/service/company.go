package service

import (
	"access-management/pkg/domain/company"
	"access-management/pkg/models"
)

type companyService struct {
	repo company.Gateway
}

func Company(repo company.Gateway) company.Service {
	return &companyService{
		repo: repo,
	}
}

func (s *companyService) Get() (models.Company, error) {
	return s.repo.Get()
}

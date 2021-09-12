package service

import (
	"access-management/pkg/models"
)

type authValidator struct{}

func (a *authValidator) ValidateUser(u *models.User) error {
	return nil
}

func (a *authValidator) ValidateCompany(u *models.Company) error {
	return nil
}

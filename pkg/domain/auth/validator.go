package auth

import "access-management/pkg/models"

type Validator interface {
	ValidateUser(u *models.User) error
	ValidateCompany(u *models.Company) error
}

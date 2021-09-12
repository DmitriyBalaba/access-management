package auth

import "access-management/pkg/models"

type Service interface {
	Signup(u *models.User) (*models.User, error)
}

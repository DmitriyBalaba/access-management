package auth

import "access-management/pkg/models"

type Gateway interface {
	Signup(u *models.User) (*models.User, error)
}

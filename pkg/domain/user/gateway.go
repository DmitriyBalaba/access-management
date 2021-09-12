package user

import (
	"access-management/pkg/models"
)

type Gateway interface {
	Get() (models.User, error)
	Create(u *models.User) error
}

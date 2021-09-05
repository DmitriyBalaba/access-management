package user

import "access-management/models"

type Repository interface {
	Get() (models.User, error)
}

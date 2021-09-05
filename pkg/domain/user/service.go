package user

import "access-management/models"

type Service interface {
	Get() (models.User, error)
}

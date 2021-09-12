package user

import (
	"access-management/pkg/models"
)

type Service interface {
	Get() (models.User, error)
}

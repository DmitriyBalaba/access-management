package company

import (
	"access-management/pkg/models"
)

type Service interface {
	Get() (models.Company, error)
}

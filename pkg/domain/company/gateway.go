package company

import (
	"access-management/pkg/models"
)

type Gateway interface {
	Get() (models.Company, error)
	Create(c *models.Company) error
}

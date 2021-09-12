package gateway

import (
	psql2 "access-management/pkg/db/psql"
	"access-management/pkg/domain/company"
	"access-management/pkg/models"
)

type companyGateway struct {
	db *psql2.DB
}

func Company(db *psql2.DB) company.Gateway {
	return &companyGateway{
		db: db,
	}
}

func (r *companyGateway) Get() (models.Company, error) {
	r.db.Print()
	return models.Company{ID: 11, Name: "Naaaa"}, nil
}

func (r *companyGateway) Create(c *models.Company) error {
	r.db.Print()
	return nil
}

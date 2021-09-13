package gateway

import (
	"access-management/pkg/db/psql"
	"access-management/pkg/domain/user"
	"access-management/pkg/models"
)

type userGateway struct {
	db *psql.DB
}

func User(db *psql.DB) user.Gateway {
	return &userGateway{
		db: db,
	}
}

func (r *userGateway) Get() (models.User, error) {
	r.db.Print()
	return models.User{
		ID:   11,
		Name: "User name",
		Company: &models.Company{
			ID:   1,
			Name: "JazzServe",
		},
	}, nil
}

func (r *userGateway) Create(u *models.User) error {
	return nil
}

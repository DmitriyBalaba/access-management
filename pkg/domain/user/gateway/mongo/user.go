package mongo

import (
	mongo2 "access-management/pkg/db/mongo"
	"access-management/pkg/domain/user"
	"access-management/pkg/models"
)

type userGateway struct {
	db *mongo2.DB
}

func User(db *mongo2.DB) user.Gateway {
	return &userGateway{
		db: db,
	}
}

func (r *userGateway) Get() (models.User, error) {
	r.db.Print()
	return models.User{
		ID:   11,
		Name: "Naaaa",
		Company: &models.Company{
			ID:   1,
			Name: "JazzServe",
		},
	}, nil
}

func (r *userGateway) Create(u *models.User) error {
	return nil
}

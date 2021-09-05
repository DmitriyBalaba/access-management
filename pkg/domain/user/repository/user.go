package repository

import (
	"access-management/models"
	"access-management/pkg/domain/db/psql"
	"access-management/pkg/domain/user"
)

type userRepository struct {
	db *psql.DB
}

func User(db *psql.DB) user.Repository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Get() (models.User, error) {
	r.db.Print()
	return models.User{ID: 11, Name: "Naaaa"}, nil
}

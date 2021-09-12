package gateway

import (
	psql2 "access-management/pkg/db/psql"
	"access-management/pkg/domain/auth"
	"access-management/pkg/domain/errors"
	"access-management/pkg/models"
)

type authGateway struct {
	db *psql2.DB
}

func Auth(db *psql2.DB) auth.Gateway {
	return &authGateway{
		db: db,
	}
}

func (r *authGateway) Signup(u *models.User) (*models.User, error) {
	if err := r.db.Create(u); err != nil {
		return nil, errors.ServerInternalError(err.Error())
	}

	if err := r.db.Create(u.Company); err != nil {
		return nil, errors.ServerInternalError(err.Error())
	}

	return u, nil
}

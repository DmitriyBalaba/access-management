//+build wireinject

package main

import (
	"access-management/models"
	"github.com/google/wire"
)

func InitApp() (*models.DB, error) {
	wire.Build(
		models.NewConfig,
		models.NewDB,
	)
	return &models.DB{}, nil  // These return values are ignored.
}

//+build wireinject

package main

import (
	"access-management/pkg/config"
	"access-management/pkg/domain/db/psql"
	"access-management/pkg/domain/user"
	"access-management/pkg/domain/user/delivery/http"
	"access-management/pkg/domain/user/repository"
	"access-management/pkg/domain/user/service"
	"access-management/pkg/server"
	"github.com/google/wire"
)

// wire.go

func InitApp() (*server.Server, error) {
	wire.Build(
		config.NewConfig,
		psql.NewDB,
		repository.User,
		service.User,
		http.User,
		user.HttpUserRouters,
		server.NewServer,
	)
	return &server.Server{}, nil // These return values are ignored.
}

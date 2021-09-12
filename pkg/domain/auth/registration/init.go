package registration

import (
	"access-management/pkg/config"
	"access-management/pkg/db/psql"
	"access-management/pkg/domain/auth/delivery/v1/http"
	"access-management/pkg/domain/auth/gateway"
	"access-management/pkg/domain/auth/service"
)

func HttpRoutes(c *config.Config) {
	http.Auth(service.Auth(gateway.Auth(psql.NewDB(c))))
}

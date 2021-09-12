package registration

import (
	"access-management/pkg/config"
	"access-management/pkg/db/psql"
	"access-management/pkg/domain/user/delivery/v1/http"
	gateway "access-management/pkg/domain/user/gateway/psql"
	"access-management/pkg/domain/user/service"
)

func HttpRoutes(c *config.Config) {
	http.User(service.User(gateway.User(psql.NewDB(c))))
}

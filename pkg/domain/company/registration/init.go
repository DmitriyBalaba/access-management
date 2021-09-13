package registration

import (
	"access-management/pkg/config"
	"access-management/pkg/db/psql"
	"access-management/pkg/domain/company/delivery/v1/http"
	"access-management/pkg/domain/company/gateway"
	"access-management/pkg/domain/company/service"
	userPsql "access-management/pkg/domain/user/gateway/psql"
)

func HttpRoutes(c *config.Config) {
	http.Company(service.Company(gateway.Company(psql.NewDB(c)), userPsql.User(psql.NewDB(c))))
}

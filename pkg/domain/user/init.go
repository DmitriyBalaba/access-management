package user

import (
	"access-management/pkg/server"
)

func HttpUserRouters(d Delivery) []server.RouteInitializer {
	return []server.RouteInitializer{InitDelivery(d)}
}

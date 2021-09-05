package user

import (
	"access-management/pkg/server"
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

type Delivery interface {
	Get(w http.ResponseWriter, r *http.Request) error
}

func InitDelivery(d Delivery) func(r *mux.Router, ctx context.Context) {
	return func(r *mux.Router, ctx context.Context) {

		server.NewRoute(r, "/users", d.Get).
			Methods(http.MethodGet)

		server.NewRoute(r, "/users/a", d.Get).
			Methods(http.MethodGet)
	}
}

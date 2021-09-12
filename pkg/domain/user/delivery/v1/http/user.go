package http

import (
	"access-management/pkg/domain/user"
	"access-management/pkg/server"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type userDelivery struct {
	userService user.Service
}

func User(userService user.Service) user.Delivery {
	d := &userDelivery{userService: userService}
	server.AddRouteInitializer(func(r *mux.Router, ctx context.Context) {

		server.NewRoute(r, "/users", d.Get).
			Methods(http.MethodGet)

		server.NewRoute(r, "/users/a", d.Get).
			Methods(http.MethodGet)
	})
	return d
}

func (u userDelivery) Get(w http.ResponseWriter, r *http.Request) error {
	out, err := u.userService.Get()
	if err != nil {
		return err
	}

	marshal, err := json.Marshal(out)
	if err != nil {
		return err
	}

	if _, err = w.Write(marshal); err != nil {
		return err
	}
	return nil
}

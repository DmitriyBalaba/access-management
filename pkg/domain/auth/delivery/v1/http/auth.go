package http

import (
	"access-management/pkg/domain/auth"
	"access-management/pkg/domain/errors"
	"access-management/pkg/models"
	"access-management/pkg/server"
	"context"
	"net/http"

	"bitbucket.org/jazzserve/webapi/web/http/payload"
	"github.com/gorilla/mux"
)

type authDelivery struct {
	auth auth.Service
}

func Auth(auth auth.Service) auth.Delivery {
	d := &authDelivery{auth: auth}

	server.AddRouteInitializer(func(r *mux.Router, ctx context.Context) {

		server.NewRoute(r, "/signup", d.Signup,
			payload.JsonBody(models.User{}, models.User{})).
			Methods(http.MethodPost)
	})
	return d
}

func (u authDelivery) Signup(w http.ResponseWriter, r *http.Request) error {
	user := payload.GetModel(r.Context()).(*models.User)
	signup, err := u.auth.Signup(user)
	if err != nil {
		errors.ResponseError(w, err)
		return nil
	}

	payload.JsonEncode(w, signup)
	return nil
}

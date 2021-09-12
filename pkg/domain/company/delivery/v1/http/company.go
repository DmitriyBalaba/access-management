package http

import (
	"access-management/pkg/domain/company"
	"access-management/pkg/server"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type companyDelivery struct {
	companyService company.Service
}

func Company(companyService company.Service) company.Delivery {
	d := &companyDelivery{companyService: companyService}

	server.AddRouteInitializer(func(r *mux.Router, ctx context.Context) {

		server.NewRoute(r, "/companies", d.Get).
			Methods(http.MethodGet)
	})
	return d
}

func (u companyDelivery) Get(w http.ResponseWriter, r *http.Request) error {
	out, err := u.companyService.Get()
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

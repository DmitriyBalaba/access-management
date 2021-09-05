package http

import (
	"access-management/pkg/domain/user"
	"encoding/json"
	"net/http"
)

type userDelivery struct {
	userService user.Service
}

func User(userService user.Service) user.Delivery {
	return &userDelivery{userService: userService}
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

func newUserDelivery(userService user.Service) *userDelivery {
	return &userDelivery{userService: userService}
}

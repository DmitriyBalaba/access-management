package auth

import "net/http"

type Delivery interface {
	Signup(w http.ResponseWriter, r *http.Request) error
}

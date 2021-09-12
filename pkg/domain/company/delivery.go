package company

import (
	"net/http"
)

type Delivery interface {
	Get(w http.ResponseWriter, r *http.Request) error
}

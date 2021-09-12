package errors

import (
	"net/http"

	"bitbucket.org/jazzserve/webapi/web/http/payload"
)

type err struct {
	Code    int `json:"code"`
	Message string
}

func (e err) Error() string {
	return e.Message
}

func NotFoundError(msg string) error {
	return &err{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func ServerInternalError(msg string) error {
	return &err{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func ResponseError(w http.ResponseWriter, e error) {
	err, ok := e.(err)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		payload.JsonEncode(w, e)
		return
	}

	w.WriteHeader(err.Code)
	payload.JsonEncode(w, err)
}

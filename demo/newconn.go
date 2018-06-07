package demo

import (
	"net/http"

	"magic.pathao.com/pinku/ebridge/helpers"
)

// NewConn of demo used to domo http handler stuff
func NewConn(msg string, w http.ResponseWriter, r *http.Request) {
	helpers.ServeJSON(w, struct {
		Msg string `json:"msg"`
	}{
		Msg: msg,
	})
}

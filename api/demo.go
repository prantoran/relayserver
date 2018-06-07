package api

import (
	"log"
	"net/http"

	"magic.pathao.com/pinku/ebridge/demo"
)

type DemoConf struct {
	Msg string `json:"msg"`
}

func DemoClient(d *DemoConf) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("starting new connection")
		demo.NewConn(d.Msg, w, r)
	}
}

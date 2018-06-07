package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// ServeJSON sends json data to the client
func ServeJSON(w http.ResponseWriter, data interface{}) {

	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(data)

	if err != nil {
		ServeInternalServerError(w)
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = io.Copy(w, buf)
	if err != nil {
		log.Println(err)
	}
}

// ServeInternalServerError send StatusInternalServerError to the client
func ServeInternalServerError(w http.ResponseWriter) {

	w.WriteHeader(http.StatusInternalServerError)
	responseJSON := map[string]interface{}{
		"error": "Internal Server Error",
	}

	ServeJSON(w, responseJSON)
}

// ServeNotFound send StatusNotFound to the client
func ServeNotFound(w http.ResponseWriter, text interface{}) {
	w.WriteHeader(http.StatusNotFound)
	responseJSON := map[string]interface{}{
		"error": text,
	}
	ServeJSON(w, responseJSON)
}

// ServeBadRequest send StatusBadRequest to the client
func ServeBadRequest(w http.ResponseWriter, msg interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	responseJSON := map[string]interface{}{
		"error": msg,
	}

	ServeJSON(w, responseJSON)
}

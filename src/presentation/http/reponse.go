package http

import (
	"encoding/json"
	netHttp "net/http"
)

func writeErrorResponse(w netHttp.ResponseWriter, err error) {
	// TODO: improve bad request error with validation details
	errRes := struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	}
	errJs, _ := json.Marshal(errRes)

	w.WriteHeader(netHttp.StatusBadRequest)
	w.Write(errJs)
}

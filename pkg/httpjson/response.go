package httpjson

import (
	"encoding/json"
	"net/http"
)

func OK(w http.ResponseWriter, resp any) {
	write(w, resp, http.StatusOK)
}

func InternalError(w http.ResponseWriter, resp any) {
	write(w, resp, http.StatusInternalServerError)
}

func write(w http.ResponseWriter, resp any, status int) {
	bytes, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(bytes)
}

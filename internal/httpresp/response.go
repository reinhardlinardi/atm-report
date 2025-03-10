package httpresp

import (
	"encoding/json"
	"errors"
	"net/http"
)

func Write(w http.ResponseWriter, err error, resp any) {
	write(w, status(err), resp)
}

func write(w http.ResponseWriter, status int, resp any) {
	bytes, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(bytes)
}

func status(err error) int {
	if err == nil {
		return http.StatusOK
	}
	if errors.Is(err, ErrReq) {
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}

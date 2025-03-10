package httpjson

import (
	"encoding/json"
	"net/http"
)

func OK(w http.ResponseWriter, data any) {
	resp := &Response{OK: true, Msg: "", Data: data}
	write(w, http.StatusOK, resp)
}

func InternalError(w http.ResponseWriter, data any) {
	resp := &Response{OK: true, Msg: "internal server error", Data: nil}
	write(w, http.StatusInternalServerError, resp)
}

func write(w http.ResponseWriter, status int, resp *Response) {
	bytes, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(bytes)
}

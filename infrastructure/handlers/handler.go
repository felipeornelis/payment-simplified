package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Handler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type BaseHandler struct{}

type response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func (bh BaseHandler) Decode(body io.Reader, payload interface{}) error {
	return json.NewDecoder(body).Decode(payload)
}

func (bh BaseHandler) Error(w http.ResponseWriter, message string, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")

	res, marshalErr := json.Marshal(response{Success: false, Data: message})
	if marshalErr != nil {
		http.Error(w, "unexpected internal error", http.StatusInternalServerError)
		log.Printf("%v", err)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(res)
}

func (bh BaseHandler) Response(w http.ResponseWriter, data interface{}, statusCode int) {
	res, err := json.Marshal(response{Success: true, Data: data})
	if err != nil {
		bh.Error(w, err.Error(), err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res)
}

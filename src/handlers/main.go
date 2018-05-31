package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data string `json:"data"`
}

func Api(w http.ResponseWriter, req *http.Request) {
	response := Response{"Api"}
	jsonResponse(response, w, http.StatusOK)
}

func Public(w http.ResponseWriter, req *http.Request) {
	response := Response{"Public"}
	jsonResponse(response, w, http.StatusOK)
}

func jsonResponse(response interface{}, w http.ResponseWriter, code int) {
	json, _ := json.Marshal(response)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

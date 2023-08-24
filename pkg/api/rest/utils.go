package rest

import (
	"encoding/json"
	"net/http"
)

const FailedToDecodeRequestBody = "failed to decode request body"

func Success(w http.ResponseWriter, T interface{}) {
	Write(w, http.StatusOK, T)
}

func BadRequest(w http.ResponseWriter, T interface{}) {
	Write(w, http.StatusBadRequest, T)
}

func Write(w http.ResponseWriter, responseStatus int, T interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(responseStatus)
	jsonObj, err := json.Marshal(T)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(jsonObj)
}

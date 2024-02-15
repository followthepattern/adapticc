package api

import (
	"encoding/json"
	"net/http"
)

const FailedToDecodeRequestBody = "failed to decode request body"

func Success(w http.ResponseWriter) {
	Write(w, http.StatusOK, nil)
}

func Created(w http.ResponseWriter) {
	Write(w, http.StatusCreated, nil)
}

func BadRequest(w http.ResponseWriter, T interface{}) {
	Write(w, http.StatusBadRequest, T)
}

func Write(w http.ResponseWriter, responseStatus int, value interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(responseStatus)

	if value == nil {
		return
	}

	jsonObj, err := json.Marshal(value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(jsonObj)
}

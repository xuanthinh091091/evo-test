package utils

import (
	"encoding/json"
	"net/http"
)

// RespondJSON write json response
func RespondJSON(w http.ResponseWriter, httpStatusCode int, object interface{}) {
	bytes, err := json.Marshal(object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpStatusCode)
	w.Write(bytes)
}

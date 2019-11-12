package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseEntity(status string, code string, took int64, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "code": code, "took": took, "message": message}
}

func Respond(respondWriter http.ResponseWriter, statusCode int, entity map[string]interface{}) {
	respondWriter.Header().Add("Content-Type", "application/json")
	respondWriter.WriteHeader(statusCode)
	json.NewEncoder(respondWriter).Encode(entity)
}

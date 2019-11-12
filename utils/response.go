package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseEntity(status string, code string, took int64, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "code": code, "took": took, "message": message}
}

func Respond(respondWriter http.ResponseWriter, entity map[string]interface{}) {
	respondWriter.Header().Add("Content-Type", "application/json")
	json.NewEncoder(respondWriter).Encode(entity)
}

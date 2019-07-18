package app

import (
	"encoding/json"
	"net/http"
)

//mapear mensaje
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "mensage": message}
}

//responder web
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

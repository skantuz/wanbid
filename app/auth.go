package app

import (
	"net/http"
	"strings"
)

var JwtAuth = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPath := r.URL.Path //direccion web
		response := make(map[string]interface{})
		contentType := r.Header.Get("Content-Type") //almacenamos el content-type en una variable
		//Verifica Tipo de Aplicacion
		if contentType != "Application/json" {
			response = Message(false, "Error en typo de Aplicacion")
			w.WriteHeader(http.StatusFailedDependency)
			Respond(w, response)
			return
		}
		//lugares sin Autorizacion
		notAuth := []string{
			"/",
		}
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenSplitted := strings.Split(r.Header.Get("Authorization"), " ")
		if len(tokenSplitted) != 2 || tokenSplitted[0] != "Bearer" {
			response = Message(false, "Token Invalido o Malformado")
			w.WriteHeader(http.StatusForbidden)
			Respond(w, response)
		}

	})
}

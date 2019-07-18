package controllers

import (
	"net/http"

	"backend/app"
)

var Index = func(w http.ResponseWriter, r *http.Request) {
	app.Respond(w, app.Message(true, "Sistema Corriendo OK"))
}

package main

import (
	r "wanbid/routes"
	"log"
	"net/http"
	"os"
)

func main() {

	router := r.NewRouter()
	port := os.Getenv("sys_port") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("Salo801105"), bcrypt.DefaultCost)
	// log.Println(string(hashedPassword))

	log.Println("Listen" + port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		log.Print(err)
	}

}

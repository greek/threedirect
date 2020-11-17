package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"apap04.com/server/config"
	"apap04.com/server/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	r.Use(config.Headers)

	r.HandleFunc("/{id}", handlers.GetLink).Methods("GET")
	r.HandleFunc("/api/id/{id}", handlers.ReturnLink).Methods("GET")
	// r.HandleFunc("/api/create", handlers.CreateLink).Methods("POST")

	var port = "3337"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	fmt.Print("listening on port " + port + "\n")
	log.Fatal(http.ListenAndServe(":"+port, r))

}

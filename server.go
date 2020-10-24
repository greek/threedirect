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
	var l = &log.Logger{
		Formatter: new(log.TextFormatter),
		Level:     log.InfoLevel,
	}

	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	r.Use(config.Headers)

	r.HandleFunc("/{id}", handlers.GetLink).Methods("GET")

	var port = "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	fmt.Print("listening on port " + port + "\n")
	l.Fatal(http.ListenAndServe(":"+port, r))

}

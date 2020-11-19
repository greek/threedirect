package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"apap04.com/server/utils"
	"github.com/gorilla/mux"
)

// Link represents a link
type Link struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte("Slug not found."))
}

// CreateLink will create a link and add it to the database.
func CreateLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// jsonfile, err := os.Open("db.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// db, _ := ioutil.ReadAll(jsonfile)

	var m Link
	err := utils.DecodeJSONBody(w, r, m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Fprintf(w, "log: ", m)
}

// GetLink will redirect the user to the link.
func GetLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["id"]

	jsonfile, err := os.Open("db.json")
	if err != nil {
		log.Fatal(err)
	}

	db, _ := ioutil.ReadAll(jsonfile)

	var m []Link
	json.Unmarshal(db, &m)

	for _, link := range m {
		if link.Id == slug {
			http.Redirect(w, r, link.Url, 301)
		}
	}

	// If we reach this, then the ID doesn't exist.
	http.Error(w, "Slug not found.", http.StatusNotFound)
}

// ReturnLink returns the details of a specific ID. Same as GetLink
// but with no redirect.
func ReturnLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	slug := vars["id"]

	jsonfile, err := os.Open("db.json")
	if err != nil {
		log.Fatal(err)
	}

	db, _ := ioutil.ReadAll(jsonfile)

	var m []Link
	json.Unmarshal(db, &m)

	for _, link := range m {
		if link.Id == slug {
			json.NewEncoder(w).Encode(link)
		}
	}

	if m == nil {
		// If we reach this, then the ID doesn't exist.
		http.Error(w, "Slug not found.", http.StatusNotFound)
	}
}

package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
}

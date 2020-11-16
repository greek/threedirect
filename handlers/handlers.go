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

// GetLink will redirect the user to the link.
func GetLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["id"]

	jsonfile, err := os.Open("db.json")
	if err != nil {
		log.Fatal(err)
	}

	db, _ := ioutil.ReadAll(jsonfile)

	var link Link
	json.Unmarshal(db, &link)

	if link.Id == slug {
		http.Redirect(w, r, link.Url, 301)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("this link doesn't exist."))
	}
}

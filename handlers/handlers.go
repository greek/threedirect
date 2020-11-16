package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Link represents a link
type Link struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

// GetLink will redirect the user to the link.
func GetLink(w http.ResponseWriter, r *http.Request) {
	var Links []Link

	Links = []Link{
		Link{Id: "H", Url: "https://apap04.com"},
		Link{Id: "goog", Url: "https://google.com"},
	}

	vars := mux.Vars(r)
	slug := vars["id"]

	// jsonfile, err := os.Open("db.json")
	// if err != nil {
	// 	log.Error(err)
	// }

	// Link, _ := ioutil.ReadAll(jsonfile)

	for _, link := range Links {
		if link.Id == slug {
			http.Redirect(w, r, link.Url, 301)
		}
	}
}

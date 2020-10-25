package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var j = []byte(`{"h":{ "source": "https://apap04.com" } }`)

// NotFound handles the fact that pages don't exist.
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Route not found"))
	log.Trace("Hi")
}

// CreateLink will either create a new link, or redirect to the link
// based on the method g.iven
func CreateLink(w http.ResponseWriter, r *http.Request) {

}

// GetLink will redirect the user to the link.
func GetLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["id"]

	var info map[string]interface{}
	if info != nil {
		fmt.Print("Route not found")
	}

	json.Unmarshal([]byte(j), &info)

	var result = info[slug].(map[string]interface{})["source"]
	http.Redirect(w, r, result.(string), 301)

}

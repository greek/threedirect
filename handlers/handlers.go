package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// NotFound handles the fact that pages don't exist.
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Route not found"))
	log.Trace("Hi")
}

// CreateLink will create a new shortlink.
func CreateLink(w http.ResponseWriter, r *http.Request) {

}

// GetLink will redirect the user to the link.
func GetLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["id"]

	jsonfile, err := os.Open("db.json")
	if err != nil {
		log.Error(err)
	}

	file, _ := ioutil.ReadAll(jsonfile)

	var info map[string]interface{}
	json.Unmarshal([]byte(file), &info)

	var result = info[slug].(map[string]interface{})["source"]
	http.Redirect(w, r, result.(string), 302)

}

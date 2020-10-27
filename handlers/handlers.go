package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Link represents the shortlink
type (
	Link struct {
		URL string `json:"source"`
	}
)

// NotFound handles the fact that pages don't exist.
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Route not found"))
	log.Trace("Hi")
}

// CreateLink will create a new shortlink.
func CreateLink(w http.ResponseWriter, r *http.Request) {
	filename := "db.json"

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error(err)
	}

	data := []Link{}
	json.Unmarshal(file, &data)
	linkStruct := &Link{
		URL: "https://apap04.com",
	}

	data = append(data, *linkStruct)

	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Error(err)
	}

	err = ioutil.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		log.Error(err)
	}
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

	var result = info["id"].(map[string]interface{})[slug]
	http.Redirect(w, r, result.(string), 302)

}

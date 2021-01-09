package apihandlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// GetLocationAndMessageHandler recieves and handle Location and Message request
func GetLocationAndMessageHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var requestBody APIRequestBody
	err = json.Unmarshal(b, &requestBody)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	requestResponse := FindShipMessageAndPosition(requestBody)
	output, err := json.Marshal(requestResponse)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

// GetLocationAndMessageSplitHandler recieves and handle Location and Message request from an specific satelite
func GetLocationAndMessageSplitHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	satelite := vars["satelite_name"]
	fmt.Println("satelite =", satelite)
}

package apihandlers

import (
	"encoding/json"
	"io/ioutil"
	"meli/quasar/satelite"
	"net/http"

	"github.com/gorilla/mux"
)

// GetLocationAndMessageHandler recieves and handle Location and Message request
func GetLocationAndMessageHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
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

	if UndeterminedResponse(requestResponse) {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)
	defer r.Body.Close()
}

// GetLocationAndMessageSplitHandler recieves and handle Location and Message request from an specific satelite
func GetLocationAndMessageSplitHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sateliteName := vars["satelite_name"]

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var requestBody APISplitRequestBody
	err = json.Unmarshal(b, &requestBody)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	requestResponse := FindShipMessageAndPositionSplit(requestBody, sateliteName)
	output, err := json.Marshal(requestResponse)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if UndeterminedResponse(requestResponse) {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)
	defer r.Body.Close()
}

// PostSateliteSplitHandler ...
func PostSateliteSplitHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sateliteName := vars["satelite_name"]

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var requestBody APISplitRequestBody
	err = json.Unmarshal(b, &requestBody)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if !ValidSateliteName(sateliteName) {
		http.NotFound(w, r)
		return
	}

	satelite.SetSateliteMessage(requestBody.Message, sateliteName)
	satelite.SetSateliteDistance(requestBody.Distance, sateliteName)

	requestResponse := "Satelite data was saved"
	output, err := json.Marshal(requestResponse)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
	defer r.Body.Close()
}

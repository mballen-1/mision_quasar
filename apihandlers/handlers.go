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
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	fmt.Println("b", b)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var requestBody APIRequestBody
	err = json.Unmarshal(b, &requestBody)
	fmt.Printf("msg = %+v", requestBody)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(requestBody)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

// GetLocationAndMessageSplitHandler recieves and handle Location and Message request for an specific satelite
func GetLocationAndMessageSplitHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	satelite := vars["satelite_name"]
	fmt.Println("satelite =", satelite)
}

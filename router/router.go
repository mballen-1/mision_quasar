package router

import (
	"meli/quasar/apihandlers"

	"github.com/gorilla/mux"
)

// DeclareRouter ..
func DeclareRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(TopSecretRoute, apihandlers.GetLocationAndMessageHandler).Methods("POST")
	router.HandleFunc(TopSecretSplitRoute, apihandlers.GetLocationAndMessageSplitHandler).Methods("GET")
	router.HandleFunc(TopSecretSplitRoute, apihandlers.PostSateliteSplitHandler).Methods("POST")
	return router
}

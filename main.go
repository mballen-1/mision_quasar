package main

import (
	"fmt"
	"meli/quasar/router"
	"meli/quasar/satelite"
	"net/http"
	"os"
)

func main() {
	satelite.ConfigureSatelites()
	router := router.DeclareRouter()
	fmt.Println("Quasar server is listening in port 5000")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.ListenAndServe(":"+port, router)
}

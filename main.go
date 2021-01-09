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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println("Quasar server is listening in port:" + port)
	http.ListenAndServe(":"+port, router)
}

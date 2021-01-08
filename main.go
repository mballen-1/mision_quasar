package main

import (
	"fmt"
	"meli/quasar/router"
	"meli/quasar/satelite"
	"net/http"
)

func main() {
	satelite.ConfigureSatelites()
	router := router.DeclareRouter()
	fmt.Println("Quasar server is listening in port 5000")
	http.ListenAndServe(":5000", router)
}

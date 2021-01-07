package main

import (
	"fmt"
	"meli/quasar/router"
	"meli/quasar/satelite"
	"net/http"
)

func configureSatelites() {
	kenobi := satelite.NewSatelite(-500, 200)
	skywalker := satelite.NewSatelite(100, -100)
	sato := satelite.NewSatelite(500, 100)

	fmt.Println("kenobi = ", kenobi)
	fmt.Println("skywalker = ", skywalker)
	fmt.Println("sato = ", sato)
}

func main() {
	configureSatelites()
	router := router.DeclareRouter()
	http.ListenAndServe(":5000", router)
	fmt.Println("Quasar server is listening in port 5000")
}

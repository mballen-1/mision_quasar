package satelite

import "fmt"

type Satelite struct {
	X float32
	Y float32
}

var Kenobi *Satelite
var Skywalker *Satelite
var Sato *Satelite

func NewSatelite(x float32, y float32) *Satelite {
	s := Satelite{X: x, Y: y}
	return &s
}

func ConfigureSatelites() {
	Kenobi = NewSatelite(-500, -200)
	Skywalker = NewSatelite(100, -100)
	Sato = NewSatelite(500, 100)

	fmt.Println("kenobi = ", Kenobi)
	fmt.Println("skywalker = ", Skywalker)
	fmt.Println("sato = ", Sato)
}

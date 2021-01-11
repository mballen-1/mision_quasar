package satelite

type Satelite struct {
	X float64
	Y float64
}

var Kenobi Satelite
var Skywalker Satelite
var Sato Satelite

func NewSatelite(x float64, y float64) Satelite {
	s := Satelite{X: x, Y: y}
	return s
}

func ConfigureSatelites() {
	Kenobi = NewSatelite(0, 1000)
	Skywalker = NewSatelite(0, 0)
	Sato = NewSatelite(-500, 500)
}

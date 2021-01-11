package satelite

// Satelite saves the coordinate of a satelite
type Satelite struct {
	X        float32
	Y        float32
	Distance float32
	Message  []string
}

// Kenobi satelite
var Kenobi Satelite

// Skywalker Satelite
var Skywalker Satelite

// Sato satelite ...
var Sato Satelite

// NewSatelite instantiates a Satelite
func NewSatelite(x float32, y float32, distance float32, message []string) Satelite {
	s := Satelite{
		X:        x,
		Y:        y,
		Distance: distance,
		Message:  message,
	}
	return s
}

// ConfigureSatelites sets coordinates of default satelites
func ConfigureSatelites() {
	Kenobi = NewSatelite(0, 1000, 500, []string{})
	Skywalker = NewSatelite(500, 500, 500, []string{})
	Sato = NewSatelite(-500, 500, 500, []string{})
}

// Kenobi = NewSatelite(0, 1000)
// 	Skywalker = NewSatelite(500, 500)
// 	Sato = NewSatelite(-500, 500)

// SetSateliteDistance ...
func SetSateliteDistance(distance float32, name string) {
	switch name {
	case "kenobi":
		{
			Kenobi.Distance = distance
		}
	case "skywalker":
		{
			Skywalker.Distance = distance
		}
	case "sato":
		{
			Sato.Distance = distance
		}
	}
}

// SetSateliteMessage ...
func SetSateliteMessage(message []string, name string) {
	switch name {
	case "kenobi":
		{
			Kenobi.Message = message
		}
	case "skywalker":
		{
			Skywalker.Message = message
		}
	case "sato":
		{
			Sato.Message = message
		}
	}
}

// GetMessages ...
func GetMessages() [][]string {
	return [][]string{
		Kenobi.Message,
		Skywalker.Message,
		Sato.Message,
	}
}

// GetDistances ...
func GetDistances() []float32 {
	return []float32{
		Kenobi.Distance,
		Skywalker.Distance,
		Sato.Distance,
	}
}

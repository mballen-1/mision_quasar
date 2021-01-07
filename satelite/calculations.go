package satelite

type satelite struct {
	x float32
	y float32
}

func NewSatelite(x float32, y float32) *satelite {
	s := satelite{x: x, y: y}
	return &s
}

func (sat satelite) GetLocation(distance float32) (x, y float32) {
	return sat.x, sat.y
}

func GetMessage(messages []string) (msg string) {
	return messages[0]
}

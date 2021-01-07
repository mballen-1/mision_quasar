package main

import "fmt"

type satellite struct {
	x float32
	y float32
}

func newSatellite(x float32, y float32) *satellite {
	s := satellite{x: x, y: y}
	return &s
}

func (sat satellite) GetLocation(distance float32) (x, y float32) {
	return sat.x, sat.y
}

func getMessage(messages []string) (msg string) {
	return messages[0]
}

func main() {
	kenobi := newSatellite(-500, 200)
	skywalker := newSatellite(100, -100)
	sato := newSatellite(500, 100)

	fmt.Println(kenobi)
	fmt.Println(skywalker)
	fmt.Println(sato)
}

package mision

import (
	"math"
	"meli/quasar/satelite"
	"strings"
)

// GetLocation := input: distancia al emisor como se recibe en cada satélite
// output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
func GetLocation(distances ...float32) (x, y float32) {
	kenobi := satelite.Kenobi
	skywalker := satelite.Skywalker
	sato := satelite.Sato

	thereIsSolution :=
		satelite.CirclesIntersect(kenobi, skywalker, distances[0], distances[1]) &&
			satelite.CirclesIntersect(kenobi, sato, distances[0], distances[2]) &&
			satelite.CirclesIntersect(skywalker, sato, distances[1], distances[2])

	if !thereIsSolution {
		return math.MaxFloat32, math.MaxFloat32
	}

	answerX, answerY, answerXPrime, answerYPrime :=
		satelite.FindIntersectionPoints(kenobi, skywalker, distances[0], distances[1])

	f64SatoX := float64(sato.X)
	f64SatoY := float64(sato.Y)

	distanceSatoToIntersectionPoint1 :=
		math.Sqrt(
			math.Pow(f64SatoX-float64(answerX), 2) +
				math.Pow(f64SatoY-float64(answerY), 2))

	distanceSatoToIntersectionPoint2 :=
		math.Sqrt(
			math.Pow(f64SatoX-float64(answerXPrime), 2) +
				math.Pow(f64SatoY-float64(answerYPrime), 2))

	if distanceSatoToIntersectionPoint1 == float64(distances[2]) {
		return float32(answerX), float32(answerY)
	}
	if distanceSatoToIntersectionPoint2 == float64(distances[2]) {
		return float32(answerXPrime), float32(answerYPrime)
	}
	return math.MaxFloat32, math.MaxFloat32
}

// GetMessage := (input): el mensaje recibido de cada satélite
// (output): el mensaje generado por el emisor
func GetMessage(messages ...[]string) (msg string) {
	if len(messages) == 0 {
		return NoMessageDetermined
	}

	sameNumberOfWordsForAllMessages := CheckSameNumberOfWordsForAllMessages(messages...)
	if !sameNumberOfWordsForAllMessages {
		return NoMessageDetermined
	}

	resultingMessage := make([]string, len(messages[0]))
	for i := range messages {
		for j := range messages[0] {
			currentMessagePart := messages[i][j]
			if currentMessagePart != "" {
				resultingMessage[j] = currentMessagePart
			}
		}
	}

	result := ""
	for v := range resultingMessage {
		result += resultingMessage[v] + " "
	}
	return strings.TrimSuffix(result, " ")
}

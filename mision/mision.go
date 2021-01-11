package mision

import (
	"fmt"
	"meli/quasar/satelite"
)

// GetLocation := input: distancia al emisor como se recibe en cada satélite
// output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
func GetLocation(distances ...float32) (x, y float32) {
	kenobi := satelite.Kenobi
	skywalker := satelite.Skywalker
	sato := satelite.Sato

	thereIsSolution :=
		satelite.CirclesTouch(kenobi, skywalker, distances[0], distances[1]) &&
			satelite.CirclesTouch(kenobi, sato, distances[0], distances[2]) &&
			satelite.CirclesTouch(skywalker, sato, distances[1], distances[2])

	if !thereIsSolution {
		return -1, -1
	}

	answerX, answerY, answerXPrime, answerYPrime := satelite.FindIntersectionPoints(kenobi, skywalker, distances[0], distances[1])

	fmt.Println(answerXPrime, answerYPrime)

	return float32(answerX), float32(answerY)
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
				j++
			}
		}
	}

	result := ""
	for v := range resultingMessage {
		result += resultingMessage[v] + " "
	}
	return result
}

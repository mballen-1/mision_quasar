package mision

// GetLocation := input: distancia al emisor como se recibe en cada satélite
// output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
func GetLocation(distances ...float32) (x, y float32) {
	return distances[0], distances[1]
}

// GetMessage := (input): el mensaje recibido de cada satélite
// (output): el mensaje generado por el emisor
func GetMessage(messages ...[]string) (msg string) {
	totalShipMessages := len(messages)
	messageLength := len(messages[0])

	messageInParts := make([]string, messageLength)
	for i := 0; i < totalShipMessages; i++ {
		for j := 0; j < messageLength; j++ {
			currentMessagePart := messages[i][j]
			if currentMessagePart != "" {
				messageInParts[j] = currentMessagePart
				j++
			}
		}
	}

	result := ""
	for v := range messageInParts {
		result += messageInParts[v] + " "
	}

	return result
}

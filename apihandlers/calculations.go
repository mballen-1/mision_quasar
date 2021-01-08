package apihandlers

import (
	"fmt"
)

var sateliteNames []string

func GetLocation(distances ...float32) (x, y float32) {
	return distances[0], distances[1]
}

func GetMessage(messages ...[]string) (msg string) {
	return messages[0][0]
}

func SolveShipMessageAndPosition(requestBody APIRequestBody) APIResponse {
	satelitesData := requestBody.Satellites
	distances, messages, sateliteNames := parseDataFromBody(satelitesData)
	fmt.Println("sateliteNames =>", sateliteNames)
	var response APIResponse
	response.Position.X, response.Position.Y = GetLocation(distances...)
	response.Message = GetMessage(messages...)
	return response
}

func parseDataFromBody(payload []SatellitePayload) (d []float32, m [][]string, n []string) {
	totalSatelites := len(payload)

	distances := make([]float32, totalSatelites)
	for i, v := range payload {
		distances[i] = v.Distance
	}

	messages := make([][]string, totalSatelites)
	for i, v := range payload {
		messages[i] = v.Message
	}

	names := make([]string, totalSatelites)
	for i, v := range payload {
		names[i] = v.Name
	}
	return distances, messages, names
}

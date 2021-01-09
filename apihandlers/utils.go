package apihandlers

import (
	"fmt"
	"meli/quasar/mision"
)

var sateliteNames []string

// FindShipMessageAndPosition parses json body and invoke mision functions to solve message and position from a ship / request
func FindShipMessageAndPosition(requestBody APIRequestBody) APIResponse {
	satelitesData := requestBody.Satellites
	distances, messages, sateliteNames := ParseDataFromBody(satelitesData)
	fmt.Println("sateliteNames =>", sateliteNames)
	fmt.Println("messages =>", messages)
	var response APIResponse
	response.Position.X, response.Position.Y = mision.GetLocation(distances...)
	response.Message = mision.GetMessage(messages...)
	return response
}

// ParseDataFromBody return distances, messages and satellite names from payload
func ParseDataFromBody(payload []SatellitePayload) (d []float32, m [][]string, n []string) {
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

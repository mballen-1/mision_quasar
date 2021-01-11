package apihandlers

import (
	"math"
	"meli/quasar/mision"
	"meli/quasar/satelite"
)

// FindShipMessageAndPosition parses json body and invoke mision functions to solve message and position from a ship / request
func FindShipMessageAndPosition(requestBody APIRequestBody) APIResponse {
	SaveDataFromBody(requestBody.Satellites)
	return AccomplishMission(satelite.GetMessages(), satelite.GetDistances())
}

// SaveDataFromBody return distances, messages and satellite names from payload
func SaveDataFromBody(payload []SatellitePayload) {
	for _, v := range payload {
		satelite.SetSateliteDistance(v.Distance, v.Name)
	}
	for _, v := range payload {
		satelite.SetSateliteMessage(v.Message, v.Name)
	}
}

// FindShipMessageAndPositionSplit parses json body and invoke mision functions to solve message and position from a ship / request
func FindShipMessageAndPositionSplit(requestBody APISplitRequestBody, sateliteName string) APIResponse {
	satelite.SetSateliteMessage(requestBody.Message, sateliteName)
	satelite.SetSateliteDistance(requestBody.Distance, sateliteName)
	return AccomplishMission(satelite.GetMessages(), satelite.GetDistances())
}

// AccomplishMission ...
func AccomplishMission(messages [][]string, distances []float32) APIResponse {
	var response APIResponse
	response.Position.X, response.Position.Y = mision.GetLocation(distances...)
	response.Message = mision.GetMessage(messages...)
	return response
}

// UndeterminedResponse ...
func UndeterminedResponse(response APIResponse) bool {
	return NoMessageFound(response) ||
		NoPositionFound(response)
}

// NoMessageFound ...
func NoMessageFound(response APIResponse) bool {
	return response.Message == mision.NoMessageDetermined
}

// NoPositionFound ...
func NoPositionFound(response APIResponse) bool {
	return response.Position.X == math.MaxFloat32 &&
		response.Position.Y == math.MaxFloat32
}

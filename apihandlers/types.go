package apihandlers

type PositionResponse struct {
	X float32
	Y float32
}

type APIResponse struct {
	Position PositionResponse
	Message  string
}

type SatellitePayload struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type APIRequestBody struct {
	Satellites []SatellitePayload `json:"satellites"`
}

type APISplitRequestBody struct {
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

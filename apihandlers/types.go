package apihandlers

type Position struct {
	X float32
	Y float32
}

type APIResponse struct {
	Position Position
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

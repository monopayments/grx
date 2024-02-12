package grx

type CreateData struct {
	Kri         string `json:"kri"`
	ReferenceID string `json:"reference_id"`
	Count       int    `json:"count"`
	Total       int    `json:"total"`
}
type CreateResponse struct {
	ReferenceID string `json:"reference_id"`
	Error       string `json:"error"`
	Message     string `json:"message"`
}

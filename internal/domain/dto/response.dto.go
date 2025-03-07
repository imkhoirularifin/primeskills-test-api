package dto

type ResponseDto struct {
	Message string      `json:"message"`
	Errors  []string    `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

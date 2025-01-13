package model

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

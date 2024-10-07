package model

type ApiResponse[T any] struct {
	ResponseCode string `json:"response_code"`
	Message      string `json:"message"`
	Data         T      `json:"data"`
}

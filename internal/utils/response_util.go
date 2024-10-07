package utils

import model "ginapi/internal/models"

func BuildResponse_[T any](status string, message string, data T) model.ApiResponse[T] {
	return model.ApiResponse[T]{
		ResponseCode: status,
		Message:      message,
		Data:         data,
	}
}

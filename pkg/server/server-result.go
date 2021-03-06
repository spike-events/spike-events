package server

import "github.com/spike-events/spike-events/bin"

var ResponseSuccess = &bin.Success{Success: true, Code: 200}

func responseError(message string) *bin.Success {
	return &bin.Success{
		Success: false,
		Code:    500,
		Message: message,
	}
}

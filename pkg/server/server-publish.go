package server

import (
	"context"
	"fmt"
	"github.com/spike-events/spike-events/bin"
)

func (s *server) Publish(ctx context.Context, message *bin.Message) (*bin.Success, error) {
	err := ctx.Err()
	if err != nil {
		return responseError(err.Error()), err
	}
	select {
	case <-ctx.Done():
		return responseError("context done"), nil
	default:
	}
	fmt.Println("publish:", message)
	err = s.sendMessage(message, true, "")
	if err != nil {
		return responseError(err.Error()), err
	}
	return ResponseSuccess, nil
}

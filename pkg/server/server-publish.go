package server

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"spike.io/bin"
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
	s.sendMessage(message, true, uuid.Nil)
	return ResponseSuccess, nil
}

package server

import (
	"context"
	"github.com/spike-events/spike-events/bin"
	"log"
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
	log.Println(">>> server publish: ", message.GetTopic(), message)
	err = s.sendMessage(message, true, "")
	log.Println("<<< server publish: ", message.GetTopic(), message)
	if err != nil {
		return responseError(err.Error()), err
	}
	return ResponseSuccess, nil
}

package server

import (
	"context"
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
	err = s.sendMessage(message, true)
	if err != nil {
		return responseError(err.Error()), err
	}
	return ResponseSuccess, nil
}


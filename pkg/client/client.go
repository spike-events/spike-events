package client

import (
	"context"
	"google.golang.org/grpc"
	"spike.io/bin"
)

type Conn interface {
	Subscribe(ctx context.Context, topic Topic) (*Subscriber, error)
	Unsubscribe(topic Topic) (*Success, error)
	Publish(topic Message) (*Success, error)
	OnDisconnect(handler func(Topic, error))
	OnError(handler func(Topic, error))
}

type srv struct {
	conn   *grpc.ClientConn
	client bin.SpikeClient

	onDisconnect func(Topic, error)
	onError      func(Topic, error)
}

func New(host string) (Conn, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := bin.NewSpikeClient(conn)
	return &srv{conn, client, nil, nil}, err
}

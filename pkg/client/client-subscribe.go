package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"spike.io/bin"
)

func (c *srv) Subscribe(ctx context.Context, topic Topic) (*Subscriber, error) {
	client := bin.NewSpikeClient(c.conn)
	stream, err := client.Subscribe(ctx, topic.ProtoMessage())
	if err != nil {
		return nil, err
	}
	return c.subscribe(ctx, topic, stream), err
}

func (c *srv) subscribe(ctx context.Context, topic Topic, stream grpc.ClientStream) *Subscriber {
	channel := make(chan *Message, 100)
	sub := &Subscriber{m: channel, connect: true, topic: topic, ctx: ctx, client: c}
	go func() {
		defer close(channel)
		for {
			select {
			case <-ctx.Done():
				stream.CloseSend()
				sub.connect = false
				return
			default:
			}
			var message bin.Message
			err := stream.RecvMsg(&message)
			if err == io.EOF {
				if c.onDisconnect != nil {
					go c.onDisconnect(fmt.Errorf("disconnect io.EOF"))
				}
				sub.connect = false
				return
			}
			if err != nil {
				if c.onError != nil {
					go c.onError(err)
				}
				continue
			}
			if message.GetOffset() == -1 {
				if c.onDisconnect != nil {
					go c.onDisconnect(fmt.Errorf("disconnect unsubscribe"))
				}
				return
			}
			channel <- protoToMessage(&message)
		}
	}()
	return sub
}

package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"spike.io/bin"
)

func (c *srv) Subscribe(ctx context.Context, topic Topic) (*Subscriber, error) {
	if len(topic.Topic) == 0 {
		return nil, fmt.Errorf("topic is required")
	}
	if len(topic.GroupId) == 0 {
		return nil, fmt.Errorf("groupId is required")
	}
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
					go c.onDisconnect(topic, fmt.Errorf("disconnect io.EOF"))
				}
				sub.connect = false
				return
			}
			if err != nil {
				if c.onError != nil {
					go c.onError(topic, err)
				}
				continue
			}
			log.Println("new message:", message)
			channel <- protoToMessage(&message)
		}
	}()
	return sub
}

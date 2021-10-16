package client

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/spike-events/spike-events/bin"
	"google.golang.org/grpc"
	"io"
	"log"
)

func (c *srv) Subscribe(ctx context.Context, topic Topic) (*Subscriber, error) {
	if len(topic.Topic) == 0 {
		return nil, fmt.Errorf("topic is required")
	}
	if len(topic.GroupId) == 0 {
		return nil, fmt.Errorf("groupId is required")
	}
	client := bin.NewSpikeClient(c.conn)
	t := topic.ProtoMessage()
	t.Id = uuid.New().String()
	stream, err := client.Subscribe(ctx, t)
	if err != nil {
		return nil, err
	}
	return c.subscribe(ctx, t, stream), err
}

func (c *srv) subscribe(ctx context.Context, topic *bin.Topic, stream grpc.ClientStream) *Subscriber {
	channel := make(chan *Message, 100)
	sub := &Subscriber{m: channel, connect: true, topic: topic, ctx: ctx, client: c}
	sub.wg.Add(1)
	connect := make(chan bool)
	defer close(connect)
	go func() {
		defer close(channel)
		defer sub.wg.Done()
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
					go c.onDisconnect(protoToTopic(topic), fmt.Errorf("disconnect io.EOF"))
				}
				sub.connect = false
				return
			}
			if err != nil {
				if c.onError != nil {
					go c.onError(protoToTopic(topic), err)
				}
				continue
			}
			log.Println(">>> client subscribe: ", message)
			if message.GetOffset() == -1 {
				if c.onDisconnect != nil {
					go c.onDisconnect(protoToTopic(topic), fmt.Errorf("disconnect io.EOF"))
				}
				sub.connect = false
				return
			}
			if message.GetTopic() == "connected" {
				connect <- true
				continue
			}
			channel <- protoToMessage(&message)
		}
	}()
	<-connect
	return sub
}

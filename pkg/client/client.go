package client

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"spike.io/bin"
)

type Conn interface {
	Subscribe(ctx context.Context, topic Topic) (*Subscriber, error)
	MonitorEvent(ctx context.Context, topic Topic) (*Subscriber, error)
	Unsubscribe(topic Topic) (*Success, error)
	Publish(topic Message) (*Success, error)
}

type srv struct {
	conn   *grpc.ClientConn
	client bin.SpikeClient
}

func New(host string) (Conn, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := bin.NewSpikeClient(conn)
	return &srv{conn, client}, err
}

func (c *srv) Unsubscribe(topic Topic) (*Success, error) {
	success, err := c.client.Unsubscribe(context.Background(), topic.ProtoMessage())
	if err != nil {
		return nil, err
	}
	return protoToSuccess(success), err
}

func (c *srv) Publish(message Message) (*Success, error) {
	success, err := c.client.Publish(context.Background(), message.ProtoMessage())
	if err != nil {
		return nil, err
	}
	return protoToSuccess(success), err
}

func (c *srv) MonitorEvent(ctx context.Context, topic Topic) (*Subscriber, error) {
	client := bin.NewSpikeClient(c.conn)
	stream, err := client.MonitorEvent(ctx, topic.ProtoMessage())
	if err != nil {
		return nil, err
	}
	return c.subscribe(ctx, topic, stream), err
}

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
				sub.connect = false
				return
			}
			if err != nil {
				log.Println(err)
				continue
			}
			channel <- protoToMessage(&message)
		}
	}()
	return sub
}

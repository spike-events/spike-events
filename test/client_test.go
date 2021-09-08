package test

import (
	"context"
	"encoding/json"
	"fmt"
	spike_io "spike.io"
	"spike.io/internal/env"
	"spike.io/pkg/client"
	"testing"
)

func TestClient(t *testing.T) {

	// start server
	connected, err := spike_io.NewServer(fmt.Sprintf(":%v", env.SpikePort))
	if err != nil {
		panic(err)
	}
	<-connected

	spikeConn, err = spike_io.NewClient(fmt.Sprintf(":%v", env.SpikePort))
	if err != nil {
		panic(err)
	}

	spikeConn.OnError(func(topic client.Topic, err error) {
		fmt.Println(topic, err)
	})
	spikeConn.OnDisconnect(func(topic client.Topic, err error) {
		fmt.Println(topic, err)
	})

	ctx := context.Background()
	sub, err := spikeConn.Subscribe(ctx, client.Topic{
		Topic:      "spike.event",
		GroupId:    "monitor",
		Offset:     0,
		Persistent: true,
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	value, _ := json.Marshal(map[string]interface{}{"ok": true})
	spikeConn.Publish(client.Message{
		Topic: "spike.event",
		Value: value,
	})

	msg := <-sub.Event()

	fmt.Println(msg)
}

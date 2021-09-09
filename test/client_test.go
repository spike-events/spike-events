package test

import (
	"context"
	"encoding/json"
	"fmt"
	spike_io "spike.io"
	"spike.io/internal/env"
	"spike.io/pkg/client"
	"sync"
	"testing"
	"time"
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
	sub1, err := spikeConn.Subscribe(ctx, client.Topic{
		Topic:      "spike.event",
		GroupId:    "monitor",
		Offset:     0,
		Persistent: true,
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	sub2, err := spikeConn.Subscribe(ctx, client.Topic{
		Topic:      "spike.event",
		GroupId:    "monitor",
		Offset:     0,
		Persistent: true,
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	sub3, err := spikeConn.Subscribe(ctx, client.Topic{
		Topic:      "spike.event",
		GroupId:    "monitor",
		Offset:     0,
		Persistent: true,
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			value, _ := json.Marshal(map[string]interface{}{"ok": true})
			wg.Add(1)
			spikeConn.Publish(client.Message{
				Topic: "spike.event",
				Value: value,
			})
			<-time.After(time.Second)
		}
	}()

	go func() {
		for msg := range sub1.Event() {
			fmt.Println("sub1:", msg)
			wg.Done()
		}
	}()

	go func() {
		for msg := range sub2.Event() {
			fmt.Println("sub2:", msg)
			wg.Done()
		}
	}()

	go func() {
		for msg := range sub3.Event() {
			fmt.Println("sub3:", msg)
			wg.Done()
		}
	}()

	wg.Wait()

	sub1.Close()
	sub2.Close()
	sub3.Close()

	sub4, err := spikeConn.Subscribe(ctx, client.Topic{
		Topic:      "spike.event",
		GroupId:    "monitor",
		Offset:     1,
		Persistent: true,
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	wg.Add(10)
	go func() {
		for msg := range sub4.Event() {
			fmt.Println("sub4:", msg)
			wg.Done()
		}
	}()

	wg.Wait()

}

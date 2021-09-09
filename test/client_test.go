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

	sub4.Close()

	for i := 0; i < 10; i++ {
		value, _ := json.Marshal(map[string]interface{}{"ok": i})
		spikeConn.Publish(client.Message{
			Topic: "spike.event",
			Value: value,
		})
	}

	sub5, err := spikeConn.Subscribe(ctx, client.Topic{
		Topic:      "spike.event",
		GroupId:    "monitor",
		Persistent: true,
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	wg.Add(10)
	go func() {
		for msg := range sub5.Event() {
			fmt.Println("sub5:", string(msg.Value))
			wg.Done()
		}
	}()

	wg.Wait()

	sub5.Close()

	sub6, err := spikeConn.Subscribe(ctx, client.Topic{
		Topic:      "spike.event",
		GroupId:    "monitor",
		Persistent: true,
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	sub7, err := spikeConn.Subscribe(ctx, client.Topic{
		Topic:      "spike.event",
		GroupId:    "monitor_2",
		Persistent: true,
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	wg.Add(2)
	go func() {
		for msg := range sub6.Event() {
			fmt.Println("sub6:", string(msg.Value))
			wg.Done()
		}
	}()
	go func() {
		for msg := range sub7.Event() {
			fmt.Println("sub7:", string(msg.Value))
			wg.Done()
		}
	}()

	spikeConn.Publish(client.Message{
		Topic:  "spike.event",
		Value:  []byte("multi channel"),
	})

	wg.Wait()

	sub7.Close()

	wg.Add(2)
	spikeConn.Publish(client.Message{
		Topic:  "spike.event",
		Value:  []byte("monitor_2 message"),
	})
	sub7, err = spikeConn.Subscribe(ctx, client.Topic{
		Topic:      "spike.event",
		GroupId:    "monitor_2",
		Persistent: true,
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	go func() {
		for msg := range sub7.Event() {
			fmt.Println("monitor_2:", string(msg.Value))
			wg.Done()
		}
	}()

	wg.Wait()
}

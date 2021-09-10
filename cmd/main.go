package main

import (
	"fmt"
	"os"
	"os/signal"
	"github.com/spike-events/spike-events/pkg/server"
	"syscall"
)

func main() {
	connected, err := server.New(fmt.Sprintf(":%v", 5672))
	if err != nil {
		panic(err)
	}
	<-connected

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
}

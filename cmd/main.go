package main

import (
	"fmt"
	"spike.io/internal/env"
	"spike.io/pkg/server"
)

func main() {
	server.New(fmt.Sprintf(":%v", env.SpikePort))
}

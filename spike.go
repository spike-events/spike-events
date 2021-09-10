package spike_io

import (
	"github.com/spike-events/spike-events/pkg/client"
	"github.com/spike-events/spike-events/pkg/server"
)

func NewClient(host string) (client.Conn, error) {
	return client.New(host)
}

func NewServer(host string) (chan bool, error) {
	return server.New(host)
}

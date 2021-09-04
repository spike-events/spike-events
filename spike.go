package spike_io

import (
	"spike.io/pkg/client"
	"spike.io/pkg/server"
)

func NewClient(host string) (client.Conn, error) {
	return client.New(host)
}

func NewServer(host string) {
	server.New(host)
}
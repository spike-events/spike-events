package server

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"spike.io/bin"
	"spike.io/pkg/server/database"
	"sync"
)

type server struct {
	subscribers map[string][]*subscribe
	m           sync.Mutex
	db          database.Service
}

func New(host string) {
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := &server{
		subscribers: make(map[string][]*subscribe),
		db:          database.New(),
	}

	s := grpc.NewServer()
	bin.RegisterSpikeServer(s, srv)

	log.Println("start server")
	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

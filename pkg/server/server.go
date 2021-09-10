package server

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"github.com/spike-events/spike-events/bin"
	"github.com/spike-events/spike-events/pkg/server/database"
	"sync"
	"time"
)

type server struct {
	subscribers               map[string][]*subscribe
	subscribersNonPersistence map[string][]*subscribe
	m                         sync.Mutex
	db                        database.Service
}

func New(host string) (chan bool, error) {
	lis, err := net.Listen("tcp", host)
	if err != nil {
		return nil, err
	}
	c := make(chan bool)
	go func() {
		srv := &server{
			subscribers: make(map[string][]*subscribe),
			db:          database.New(),
		}

		s := grpc.NewServer()
		bin.RegisterSpikeServer(s, srv)

		log.Println("start server")

		// and start...c
		go func() {
			<-time.After(time.Millisecond * 500)
			c <- true
		}()
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return c, nil

}

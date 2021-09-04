package env

import (
	"os"
	"strconv"
)

var SpikePort = 5672
var ConnectionString = os.Getenv("DB_CONN")
var MaxConnections = 5

func init() {
	port, err := strconv.Atoi(os.Getenv("SPIKE_PORT"))
	if err == nil {
		SpikePort = port
	}
	maxConn, err := strconv.Atoi(os.Getenv("DB_MAX_CONN"))
	if err != nil || maxConn > 0 {
		MaxConnections = maxConn
	}
}

package maggieQ

import "net"
import "fmt"
import "sync"
import "math/rand"
import "bufio"
import "time"

// global connection store
var Connections ConnectionMap

type ConnectionMap struct {
	ConnectionStore map[int64]*Connection
	sync.Mutex
}

type Connection struct{
	ID int64
	Channels map[int64]*Channel
	TCPConnection *net.Conn
}


var port int

func ServeForever() error{
	portStr := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", portStr)
	if err != nil {
		// handle error
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			return err
		}
		go NewConnection(conn)
	}
}


func NewConnection(conn net.Conn) error {
	// Get a random ID for this connection
	Connections.Lock()
	for {
		randID := rand.Int63()
		_, ok := Connections.ConnectionStore[randID]
		if ok {
			continue
		} else {
			connection := &Connection{
				ID: randID,
				Channels: make(map[int64]*Channel),
				TCPConnection: &conn,
			}
			Connections.ConnectionStore[randID] = connection
			break
		}
	}
	Connections.Unlock()
	// Read from the connection
	conn.SetReadDeadline(time.Second * 30)
	r := bufio.NewReader(conn)
	for {
		//handle messages sent from client
	}
}


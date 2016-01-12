package maggieQ

import "net"
import "fmt"
import "sync"
import "math/rand"
import "bufio"
import "time"

// global connection store
var Connections ConnectionMap

const MAX_CHANNELS_PER_CONNECTION = 65535

type ConnectionMap struct {
	ConnectionStore map[int64]*Connection
	sync.Mutex
}

type Connection struct{
	ID int64
	sync.Mutex
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

func (conn *Connection) OpenChannel() (Channel, error) {
	channelCount := len(conn.Channels)
	if channelCount >= MAX_CHANNELS_PER_CONNECTION {
		return nil, ERROR_MAX_CHANNEL_REACHED
	}
	// Create a new channel
	var channel Channel
	conn.Lock()
	for {
		randID := rand.Int63()
		_, ok := conn.Channels[randID]
		if ok {
			continue
		} else {
			channel = &Channel{
				ID: randID,
				Consumers: make(map[string]*Consumer),
			}
			conn.Channels[randID] = channel
			break
		}
	}
	conn.Unlock()
	// TODO: Start a channel handler goroutine

	return channel, nil
}
package model

import "net"

type SocketServerTask struct {
	Address     string
	ProcessFunc func(conn net.Conn)
}

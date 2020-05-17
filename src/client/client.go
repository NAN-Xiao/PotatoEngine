package client

import (
	"net"
	"potatoengine/src/connection"
)

type Client struct {
	_conn connection.Connnetion
}

func NewClient(conn *net.TCPConn) *Client {

	_conn:=connection.NewConnection(conn)
	client :=&Client{_conn:*_conn}
	return  client
}


}
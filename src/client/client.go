package client

import (
	"net"
	"potatoengine/src/connection"
	"potatoengine/src/message"
)

type Client struct {
	_cID  uint32
	_conn connection.Connnetion
}

func (cl *Client) Send(msg *message.Messsage) {
	if msg == nil {
		return
	}
	cl._conn.SendMessage(msg)
}
func NewClient(conn *net.TCPConn) *Client {
	_conn := connection.NewConnection(conn)
	client := &Client{_conn: *_conn}
	return client
}

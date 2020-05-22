package client

import (
	"net"
	"potatoengine/src/connection"
	"potatoengine/src/message"
)

type Client struct {
	_cID  uint32
	_conn *connection.Connnetion
}

func OnConnection()  {
}


func (cl *Client) Send(msg *message.Messsage) {
	if msg == nil {
		return
	}
	cl._conn.Write(msg)
}

func (cl *Client) Receive() {

}

func NewClient(conn *net.TCPConn) *Client {
	_tconn := connection.NewConnection(conn)
	client := &Client{_conn: _tconn}
	return client
}

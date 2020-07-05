package client

import (
	"net"
	"potatoengine/src/agent"
	"potatoengine/src/dispatcher"
)

type Client struct {
	UID         int32
	Agent       *agent.Agent
	dispatch    *dispatcher.Dispatcher
	ReadChanel  chan interface{}
	WriteChanel chan interface{}
	_conn       net.Conn
}

func (this *Client) WriteToChanle(msg interface{}) {
	this.WriteChanel <- msg
}
func (this *Client) ReadToChanle() interface{} {
	return <-this.ReadChanel
}
func NewClient(conn net.Conn) *Client {
	client := &Client{
		UID:         0,
		dispatch:    &dispatcher.Dispatcher{},
		_conn:       conn,
		ReadChanel:  make(chan interface{}, 128),
		WriteChanel: make(chan interface{}),
	}
	return client
}

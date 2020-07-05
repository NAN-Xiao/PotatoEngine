package client

import (
	"net"
	"potatoengine/src/agent"
	"potatoengine/src/dispatcher"
)

type Client struct {
	UID      int32
	Agent    *agent.Agent
	dispatch *dispatcher.Dispatcher
	// ReadChanel  chan interface{}
	MsgChanel chan interface{}
	_conn     net.Conn
}

func (this *Client) WriteToChanle(msg interface{}) {
	this.MsgChanel <- msg
}

//派发消息
func (this *Client) DispatchMsg() {
	for {
		m, ok := <-this.MsgChanel
		if ok == false {
			continue
		}
		this.dispatch.Dispach(m)
	}
}

func NewClient(conn net.Conn) *Client {
	client := &Client{
		UID:       0,
		dispatch:  &dispatcher.Dispatcher{},
		_conn:     conn,
		MsgChanel: make(chan interface{}, 128),
	}
	go client.DispatchMsg()
	return client
}

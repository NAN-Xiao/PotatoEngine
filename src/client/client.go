package client

import (
	"net"
	"potatoengine/src/agent"
	"potatoengine/src/dispatcher"
	"potatoengine/src/netmessage"
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
		msg := netmessage.PackMessagePackage(this.UID, this.Agent.GetPlayerID(), m)
		if msg == nil {
			continue
		}
		this.dispatch.Dispatch(msg)
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

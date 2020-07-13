package client

import (
	"net"
	"potatoengine/src/account"
	"potatoengine/src/agent"
	"potatoengine/src/dispatcher"
	"potatoengine/src/netmessage"
)

type Client struct {
	UID      int32
	dispatch *dispatcher.Dispatcher
	// ReadChanel  chan interface{}
	MsgChanel chan interface{}
	Account *account.Account
	Agent    *agent.Agent
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
		var pid int32 = 0
		if this.Agent != nil {
			pid = this.Agent.GetPlayerID()
		}
		msg := netmessage.PackMessagePackage(this.UID, pid, m)
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

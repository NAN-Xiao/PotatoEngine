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
	RecevieChan  chan interface{}
	SendChan chan interface{}
	Account    *account.Account
	Agent      *agent.Agent
	_conn      net.Conn
}
//接受断言好的消息
func (this *Client) Recevie(msg interface{}) {

	this.RecevieChan <- msg
}
func (this *Client)Send(msg interface{})  {
	if msg==nil||this.SendChan==nil{
		return
	}
	this.SendChan<-msg
}
//派发消息
func (this *Client) DispatchMsg() {
	for {
		m, ok := <-this.RecevieChan
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
//发送网络端消息到客户端
func (this *Client) SendToNet()  {
	if this._conn == nil {
		return
	}
	for {
		if m, ok := <-this.SendChan; ok == true {
			data, err := netmessage.PackageNetMessage(m)
			if err == nil {
				this._conn.Write(data)
			}
		} else {
			break
		}
	}
}
func NewClient(conn net.Conn) *Client {
	client := &Client{
		UID:        0,
		dispatch:   &dispatcher.Dispatcher{},
		_conn:      conn,
		RecevieChan:  make(chan interface{}, 128),
		SendChan: make(chan interface{}, 128),
	}
	go client.DispatchMsg()
	return client
}

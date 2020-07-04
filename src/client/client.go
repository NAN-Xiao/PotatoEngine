package client

import (
	"potatoengine/src/agent"
	"potatoengine/src/connection"
	"potatoengine/src/netmessage"
)

type Client struct {
	agent 	*agent.Agent
	readChanel chan interface{}
	writeChanel chan interface{}
	_conn      connection.IConn
}
func (this *Client)WriteToChanle(msg interface){
	writeChanel->msg
}
func(this *Client)ReadToChanle()interface{}{
	return<-ReadChanel
}
func NewClient(conn connection.IConn) *Client {
	client := &Client{
		UserID:     0,
		PlayerID:   0,
		_conn:      conn,
		Readchanel: make(chan interface{}, 128),
	}
	return client
}
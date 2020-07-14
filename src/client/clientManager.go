package client

import (
	"potatoengine/src/connection"
)

type ClientMap struct {
	Clients []Client
	ClientIndex int32
}
func (this *ClientMap)AddClient(conn connection.IConn)  {
	this.ClientIndex+=1
	cl:=Client{
		ClientID:    this.ClientIndex,
		//RecevieChan: make(chan interface{},128),
		SendChan:    make(chan interface{},128),
		Account:     nil,
		Agent:       nil,
		Conn: conn,
	}
	this.Clients = append(this.Clients, cl)
	cl.Recevie()
}
//删除客户端
func (this *ClientMap)RemoveClient(cl Client)  {
	if this.Clients==nil|| len(this.Clients)<=0{
		return
	}
	for i:=range this.Clients{
		if this.Clients[i]==cl{
			this.Clients=append(this.Clients[:i],this.Clients[i+1:]...)
		}
	}
}
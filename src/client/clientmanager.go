package client

import (
	"fmt"
	"potatoengine/src/netmessage"
)

var Clients []*Client

//当有连接。添加持有的客户端
func (this *ClientMgr) AddClient(cl *Client) error {

	for C := range Clients {
		if c == cl {
			
			return fmt.Errorf("this client contains")
		}
	}
	Clients=append Clients cl
	return nil
}
//删除持有的客户端
func (this *ClientMgr) RemoveCLient(cl *Client) {
	
}
//广播消息
func  BroadcastMessage(msg *netmessage.ServerMsgPackage) {

	if mgr._clients == nil {
		return
	}
	for cl:= range Clients {
		cl.WriteToChanle(msg)
	}
}

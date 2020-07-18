package client

import (
	"fmt"
	"potatoengine/src/account"
	"potatoengine/src/agent"
	"potatoengine/src/connection"
	"potatoengine/src/logService"
	"potatoengine/src/netmessage"
)

type Client struct {
	ConnID int32
	SendChan chan interface{}
	Conn     connection.IConn
	Account  *account.Account
	Agent    agent.Agent
}

//接受断言好的消息 开始接受发送线程
func (this *Client) Recevie() {
	if this.Conn == nil {
		return
	}
	go this.Conn.Receive()
	go this.Conn.Send()
}

//发送网络端消息到客户端
func (this *Client) SendToNet() {
	if this.Conn == nil||this.Conn.IsClosed() {
		logService.LogError(fmt.Sprintf("connect is closed : client id :%d",this.ConnID))
		return
	}
	for {
		if m, ok := <-this.SendChan; ok == true {
			data, err := netmessage.PackageNetMessage(m)
			if err == nil {
				this.Conn.WriteMsg(data)
			}
		} else {
			break
		}
	}
}

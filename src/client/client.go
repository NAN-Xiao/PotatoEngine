package client

import (
	"potatoengine/src/account"
	"potatoengine/src/agent"
	"potatoengine/src/connection"
	"potatoengine/src/netmessage"
)

type Client struct {
	ClientID int32
	SendChan chan interface{}
	Account  *account.Account
	Agent    *agent.Agent
	Conn     connection.IConn
}

//接受断言好的消息 写入
func (this *Client) Recevie() {
	if this.Conn == nil {
		return
	}
	go this.Conn.Receive()
}
func (this *Client) Send(msg interface{}) {
	if msg == nil || this.SendChan == nil {
		return
	}
	this.SendChan <- msg
}

//todo 读取接受的缓冲消息并派发到对应处理模块
func (this *Client) DisPatchMsg() {
	que := this.Conn.GetMsgQue()
	if que == nil || len(que) <= 0 {
		return
	}
	for {
		msg, ok := <-que
		if ok {
			//todo
			id,err:=netmessage.GetServerMsgID(msg)
			if err!=nil{
				continue
			}
			//登录相关消息
			if id>10000&&id<20000{
				
			}else if id>20000&&id<30000 {

			}
		}
	}
}

//发送网络端消息到客户端
func (this *Client) SendToNet() {
	if this.Conn == nil {
		return
	}
	for {
		if m, ok := <-this.SendChan; ok == true {
			data, err := netmessage.PackageNetMessage(m)
			if err == nil {
				this.Conn.Write(data)
			}
		} else {
			break
		}
	}
}

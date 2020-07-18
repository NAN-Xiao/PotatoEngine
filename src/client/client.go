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

//接受断言好的消息 写入
func (this *Client) Recevie() {
	if this.Conn == nil {
		return
	}
	go this.Conn.Receive()
	go this.Conn.()
}


////todo 读取接受的缓冲消息并派发到对应处理模块
//func (this *Client) DisPatchMsg() {
//	que := this.Conn.GetMsgQue()
//	if que == nil || len(que) <= 0 {
//		return
//	}
//	for {
//		msg, ok := <-que
//		if ok {
//			//todo
//			id,err:=netmessage.GetServerMsgID(msg)
//			if err!=nil{
//				continue
//			}
//			//登录相关消息
//			if id>10000&&id<20000{
//
//			}else if id>20000&&id<30000 {
//
//			}
//		}
//	}
//}

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

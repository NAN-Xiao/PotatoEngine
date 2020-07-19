package client

import (
	"fmt"
	"net"
	"potatoengine/src/account"
	"potatoengine/src/agent"
	"potatoengine/src/connection"
	"potatoengine/src/logService"
	"potatoengine/src/netmessage"
)

type Client struct {
	SendChan chan interface{}
	Conn     connection.IConn
	Account  *account.Account
	Agent    *agent.Agent
	Conn_id connection.ConnID
}

//开始接受发送线程
func (this *Client) Connect() {
	if this.Conn == nil {
		return
	}
	go this.Conn.Receive()
	go this.Conn.Send()
}
func (this *Client) DispatchMsg() {
	for {
		if this.Conn.IsClosed() {
			this.Delete()
			break
		}
		if this.Account==nil{
			this.Account=account.NewAccount(this)
		}
		msg:=this.Conn.Read()
		if msg!=nil{
			id,err:=netmessage.GetServerMsgID(msg)
			if err!=nil{
				continue
			}
			//if id>0&&id<
		}
	}
}

//发送网络端消息到客户端
func (this *Client) Send() {
	if this.Conn == nil || this.Conn.IsClosed() {
		logService.LogError(fmt.Sprintf("connect is closed : client id :%d", this.Conn_id.Get()))
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

//从space移除account和angent 并移除连接
func (this *Client) Delete() {
	//关闭连接
	addr := this.Conn.GetRemoteAddr().String()
	if !this.Conn.IsClosed() {
		this.Conn.Close()
	}
	//删除account
	if this.Account != nil {
		sp := this.Account.Entity.GetCurrentSpace()
		this.Account.Entity.LeaveSpace(sp.SpaceID)
		this.Account = nil
	}
	//删除angent
	if this.Agent != nil {
		sp := this.Agent.Entity.GetCurrentSpace()
		this.Agent.Entity.LeaveSpace(sp.SpaceID)
		this.Agent = nil
	}
	DeleteClient(this)
	logService.Log(fmt.Sprintf("Close CLient Connect,Remoteinfo::%s", addr))
}

func NewClient(con net.Conn)  *Client{
	tcp,ok:=con.(*net.TCPConn)
	if !ok{
		logService.LogError(fmt.Sprintf("creat new client fail,because this net connect is not  tcp"))
		return nil
	}
	tcpcon :=&connection.TcpConnect{
		Conn:    tcp,
		ReceiveChan: make(chan interface{}, 128),
		ConnID:      0,
		SendChan:    nil,
	}
	cl:=&Client{
		ConnID:   0,
		SendChan: make(chan interface{},128),
		Conn:     tcpcon,
		Account:  nil,
		Agent:    nil,
	}
	return cl
}
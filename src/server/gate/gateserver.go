package gate

import (
	"fmt"
	"net"
	"potatoengine/src/client"
	"potatoengine/src/message"
	"potatoengine/src/server"
	"potatoengine/src/space"
)

type GateServer struct {
	server.BaseServer
}

func (this *GateServer) Initialize() {

}

//内部调用
func (this *GateServer) Begin() {

	//this.RunSpace()

	addr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:8999")
	if err != nil {
		fmt.Println("loginserver start error from resolve addr")
		return
	}
	lisenter, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println("listener err")
		return
	}
	go func() {
		for {
			//fmt.Println("listen client connect")
			tcpConn, err := lisenter.AcceptTCP() //阻塞，当有客户端连接时，才会运行下面
			if err != nil {
				//fmt.Println("tcpListener error :", err)
				continue
			}
			cl := client.NewClient(tcpConn)
			client.GetClientMgr().AddClient(cl)
		}
	}()
	fmt.Println("GaterServer Started")
}

func (this *GateServer) Run() {
	this.Begin()
}
func (this *GateServer) Stop() {

}

//广播消息
func (this *GateServer) BroadcastMessage(msg *message.MsgPackage) {

	if msg == nil {
		return
	}
	m := msg.GetMessage()
	client.GetClientMgr().BroadcastMessage(m)
}

//todo
//可以通过一个全局接口方法根据type创建不同的server
///新建gateserver
func NewGateServer() *GateServer {
	sr := &GateServer{struct {
		Listener *net.TCPListener
		Spaces   map[string]space.ISpace
	}{Listener: nil, Spaces: make(map[string]space.ISpace)}}
	return sr
}

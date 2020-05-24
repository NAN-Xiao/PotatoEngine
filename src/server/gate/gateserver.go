package gate

import (
	"fmt"
	"net"
	"potatoengine/src/client"
	"potatoengine/src/space"
)

type GateServer struct {
	_listener *net.TCPListener
	_spaces   map[string]space.ISpace
}

//为当前服务注册space
func (this *GateServer) RegisterSpace(sp space.ISpace) {
	if sp == nil {
		return
	}
	name := sp.GetName()
	_, ok := this._spaces[name]
	if ok == false {
		fmt.Println("")
	}
	this._spaces[name] = sp
}

func (this *GateServer) Initialize() {

}

//内部调用
func (this *GateServer) Begin() {

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
	for sp := range this._spaces {
		go func() {
			for{
				this._spaces[sp].Process()
			}
		}()

	}
	//defer lisenter.Close()
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
			//go client.GetClientMgr().
		}
	}()
	fmt.Println("GaterServer Started")
}

func (this *GateServer) Start() {
	this.Begin()
}
func (this *GateServer) Stop() {

}

func NewGateServer() *GateServer {
	ser := &GateServer{
		_listener: nil,
		_spaces:   make(map[string] space.ISpace),
	}
	return ser
}

package server

import (
	"fmt"
	"net"
	"potatoengine/src/client"
	"potatoengine/src/connection"
	"potatoengine/src/globleTimer"
	"potatoengine/src/space"
)

type BaseServer struct {
	AllClient client.ClientMap
	SpacesMap map[string]space.ISpace
	ConType   connection.ConnType
	Name      E_ServerNames
}

func (this *BaseServer) RegisterSpace(sp space.ISpace) {
	if sp == nil {
		return
	}
	name := sp.GetName()
	_, ok := this.SpacesMap[name]
	if ok {
		fmt.Printf("have current space::%s \n", name)
		return
	}

	this.SpacesMap[name] = sp
	fmt.Printf("RegisterSpace::%s \n", name)
}
func (this *BaseServer) Stop() {
	//todo 断开所有的客户端链接 卸载所有的space
}

//启动服务器
func (this *BaseServer) Run() {
	//todo 启动tick的携程 这里开启了新的线程来更新tick，主要目的是全局唯一的tick
	globleTimer.Tick()
	//启动监听 top｜http
	if this.ConType == connection.ETcp {
		go this.ListenTcp()
	} else {
		go this.ListenHttp()
	}
}

//阻塞监听tcp。当有链接创建tcp链接客户端 并添加到服务器持有到客户端链接队列
func (this *BaseServer) ListenTcp() {
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:9000")
	if err != nil {
		println(err)
		return
	}
	lisener, err := net.ListenTCP("tcp", addr)

	for {
		c, err := lisener.AcceptTCP()
		if err != nil {
			println(err)
			return
		}
		con := &connection.TcpConnect{
			Conn:    c,
			MsgChan: make(chan interface{}, 128),
		}
		cl:=client.Client{
			SendChan:    make(chan interface{},128),
			Account:     nil,
			Agent:       nil,
			Conn: con,
		}
		this.AllClient.AddClient(cl)
		cl.Recevie()
	}
}

//todo 监听http
func (this *BaseServer) ListenHttp() {

}

func (this *BaseServer) SpaceRun() bool {
	if this.SpacesMap == nil || len(this.SpacesMap) <= 0 {
		fmt.Printf("this server have any space ::%d \n", len(this.SpacesMap))
		return false
	}
	for s := range this.SpacesMap {
		sp := this.SpacesMap[s]
		if sp == nil {
			continue
		}
		fmt.Printf("RunSpace Name::(%s)>>\n", sp.GetName())
		go sp.Process()
	}
	return true
}

func NewServer(srname E_ServerNames, connType connection.ConnType) *BaseServer {
	sr := &BaseServer{
		SpacesMap: make(map[string]space.ISpace),
		Name:      srname,
		ConType:   connType,
		AllClient: client.ClientMap{
			Clients:     make([]client.Client, 0),
			ClientIndex: 0,
		},
	}
	return sr
}

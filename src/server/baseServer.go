package server

import (
	"fmt"
	"potatoengine/src/common"
	"potatoengine/src/logService"
	"potatoengine/src/netWork/connect"
	"potatoengine/src/netWork/listener"
	"potatoengine/src/space"
)

type BaseServer struct {
	SpacesMap map[string]space.ISpace
	ConType   connect.ConnType
	Name      E_ServerNames
	Listener  listener.IListener
}

//注册当前server的space
func (this *BaseServer) RegisterSpace(sp space.ISpace) {
	if sp == nil {
		return
	}
	spaceinfo:=sp.GetSpace()
	if spaceinfo==nil{
		return
	}
	name :=spaceinfo.Spacename
	_, ok := this.SpacesMap[name]
	if ok {
		fmt.Printf("have current space::%s \n", name)
		return
	}
	this.SpacesMap[name] = sp
	space.AddSpace(sp)
	fmt.Printf("RegisterSpace::%s \n", name)
}

//停止serve
func (this *BaseServer) Stop() {
	//todo 断开所有的客户端链接 卸载所有的space
}

//启动服务器
func (this *BaseServer) Run() {
	//todo 启动tick的携程 这里开启了新的线程来更新tick，主要目的是全局唯一的tick
	common.Tick()
	//启动监听 top｜http
	switch this.ConType {
		case connect.ETcp:
			this.Listener = listener.NewTcpListener("tcp", "0.0.0.0:8999")
			this.Listener.Listen()
		case connect.EHttp:
			logService.LogError("gateserver cant use http netWork")
	}
	this.SpaceRun()
}

//启动space 并注册sp中的tik函数
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
		common.RegiestTick(sp.Tick)
		sp.OnStart()
		go sp.Process()
		logService.Log(fmt.Sprintf(" space(name::%s)is run",sp.GetSpace().Spacename))
	}
	return true
}

func NewServer(srname E_ServerNames, connType connect.ConnType) *BaseServer {
	sr := &BaseServer{
		SpacesMap: make(map[string]space.ISpace),
		Name:      srname,
		ConType:   connType,
	}
	return sr
}

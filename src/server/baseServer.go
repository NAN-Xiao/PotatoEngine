package server

import (
	"fmt"
	"potatoengine/src/netWork"
	"potatoengine/src/globleTimer"
	"potatoengine/src/logService"
	"potatoengine/src/space"
)

type BaseServer struct {
	SpacesMap map[string]space.ISpace
	ConType   netWork.ConnType
	Name      E_ServerNames
	Listener  netWork.IListener
}

//注册当前server的space
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
	spaceMgr.AddSpace(sp)
	fmt.Printf("RegisterSpace::%s \n", name)
}

//停止serve
func (this *BaseServer) Stop() {
	//todo 断开所有的客户端链接 卸载所有的space
}

//启动服务器
func (this *BaseServer) Run() {
	//todo 启动tick的携程 这里开启了新的线程来更新tick，主要目的是全局唯一的tick
	globleTimer.Tick()
	//启动监听 top｜http
	switch this.ConType {
		case netWork.ETcp:
			this.Listener = netWork.NewTcpListener("tcp", "0.0.0.0:8999")
			this.Listener.Listen()
		case netWork.EHttp:
			logService.LogError("gameserver cant use http netWork")
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
		globleTimer.RegiestTick(sp.Tick)
		sp.OnStart()
		go sp.Process()
		logService.Log(fmt.Sprintf(" space(%s)is run",sp.GetName()))
	}
	return true
}

func NewServer(srname E_ServerNames, connType netWork.ConnType) *BaseServer {
	sr := &BaseServer{
		SpacesMap: make(map[string]space.ISpace),
		Name:      srname,
		ConType:   connType,
	}
	return sr
}

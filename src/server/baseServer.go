package server

import (
	"fmt"
	"potatoengine/src/connection"
	"potatoengine/src/space"
)

type BaseServer struct {
	Conn   connection.IConn
	SpacesMap map[string]space.ISpace
	Name   E_ServerNames
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
}

func (this *BaseServer) Run() {
	//启动space
	//if !this.SpaceRun() {
	//	fmt.Println("game server start space run is fail")
	//	return
	//}
	//启动监听
	this.Conn.Listen()
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


package server

import (
	"fmt"
	"net"
	"potatoengine/src/space"
)

type BaseServer struct {
	Listener *net.TCPListener
	Spaces   map[string]space.ISpace
}

func (this *BaseServer) RunSpace() bool {
	sp := this.Spaces
	if sp == nil {

		return false
	}
	for s := range sp {
		go sp[s].Process()
	}
	return true
}

//为当前服务注册space
func (this *BaseServer) RegisterSpace(sp space.ISpace) {
	if sp == nil {
		return
	}
	name := sp.GetName()
	_, ok := this.Spaces[name]
	if ok == false {
		fmt.Println("")
	}
	this.Spaces[name] = sp
}

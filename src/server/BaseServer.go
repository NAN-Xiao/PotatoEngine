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

func (this *BaseServer) RegisterSpace(sp space.ISpace) {
	if sp == nil {
		return
	}
	name := sp.GetName()
	_, ok := this.Spaces[name]
	if ok {
		return
	}
	this.Spaces[name] = sp
	fmt.Printf("SpaceAdded::%s \n", name)
	fmt.Printf("SpacesLen::%d \n", len(this.Spaces))
}
func (this *BaseServer) Stop() {
}

func (this *BaseServer) Run(){
	ok := this.SpaceRun()
	if ok == false {
		return
	}
}
func (this *BaseServer) SpaceRun() bool {
	if this.Spaces == nil || len(this.Spaces) <= 0 {
		fmt.Printf("this server have any space ::%d \n", len(this.Spaces))
		return false
	}
	for s := range this.Spaces {
		sp := this.Spaces[s]
		if sp == nil {
			continue
		}
		fmt.Printf("RunSpace::getname(%s)>>\n", sp.GetName())
		go sp.Process()
	}
	//fmt.Println("runspace")
	return true
}
func NewServer() *BaseServer {
	sv:=&BaseServer{

	}
	return sv
}



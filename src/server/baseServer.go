package server

import (
	"fmt"
	"potatoengine/src/space"
)

type BaseServer struct {
	Spaces   map[string]space.ISpace
	Name E_ServerNames
}

func (this *BaseServer) RegisterSpace(sp space.ISpace) {
	if sp == nil {
		return
	}
	name := sp.GetName()
	_, ok := this.Spaces[name]
	if ok {
		fmt.Printf("have current space::%s \n", name)
		return
	}
	this.Spaces[name] = sp
	fmt.Printf("SpaceAdded::%s \n", name)
	//fmt.Printf("SpacesLen::%d \n", len(this.Spaces))
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
		fmt.Printf("RunSpace Name::(%s)>>\n", sp.GetName())
		go sp.Process()
	}
	return true
}




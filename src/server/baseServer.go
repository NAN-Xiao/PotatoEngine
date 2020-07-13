package server

import (
	"fmt"
	"potatoengine/src/connection"
	"potatoengine/src/space"
	"time"
)

type BaseServer struct {
	Conn      connection.IConn
	SpacesMap map[string]space.ISpace
	Name      E_ServerNames
	tick      *time.Ticker
	tickfn    []func()
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
	this.tickfn = append(this.tickfn, sp.Tick)
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

	//todo 启动tick的携程 这里开启了新的线程来更新tick，主要目的是全局唯一的tick
	//以后希望能和space的tick放到一起。
	go func() {
		//println("start tick")
		for {
			select {
			case <-this.tick.C:
				ln := len(this.tickfn)
				if ln <= 0 {
					continue
				}
				for i := 0; i < ln; i++ {
					fn := this.tickfn[i]
					fn()
				}
			}
		}
	}()
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

func NewServer(srname E_ServerNames, connType connection.ConnType) *BaseServer {
	sr := &BaseServer{
		SpacesMap: make(map[string]space.ISpace),
		Name:      E_Game,
	}
	if connType == connection.ETcp {
		sr.Conn = &connection.TcpConnect{}

	}
	if connType == connection.EHttp {
		sr.Conn = &connection.HttpConnect{}
	}
	sr.tick = time.NewTicker(time.Second / 2)
	sr.tickfn = make([]func(), 0)

	return sr
}

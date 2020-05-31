package main

import (
	"fmt"
	"net"
	"potatoengine/src/agent"
	"potatoengine/src/client"
	"potatoengine/src/message"
	"potatoengine/src/space"
)

type LoginSpace struct {
	space.BaseSpace
}

func (this *LoginSpace) Process() {
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
}

//agent 进入场景
func (this *LoginSpace) LeaveSpace(ag *agent.Agent) {
	v, ok := this.Agents[ag.GetPlayerID()]
	if ok {
		v.OnLeaveSpace()
		delete(this.Agents, v.GetPlayerID())
	}
}

//agent退出场景
func (this *LoginSpace) EnterSpace(ag *agent.Agent) {
	_, ok := this.Agents[ag.GetPlayerID()]
	if ok {
		return
	}
	this.Agents[ag.GetPlayerID()] = ag
	ag.OnLeaveSpace()
}

func (this *LoginSpace) GetName() string {
	return this.Spacename
}

func NewLoginSpace(name string) space.ISpace {
	sp := &LoginSpace{space.BaseSpace{
		SpaceID:    0,
		Spacename:  name,
		Agents:     make(map[uint32]*agent.Agent),
		Spacechanl: make(chan *message.MsgPackage, 100),
	}}
	return sp
}

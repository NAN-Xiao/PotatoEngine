package main

import (
	"fmt"
	"potatoengine/src/agent"
	"potatoengine/src/message"
	"potatoengine/src/space"
)

type LoginSpace struct {
	space.BaseSpace
}

func (this *LoginSpace) Process() {
	for {
		value, ok := <-this.Spacechanl
		if ok == false {
			continue
		}
		msgid := value.GetMsgID()
		_, have := message.NetMessageType[msgid]
		if have == false {
			fmt.Printf("%s space process mesg error bac not register %d msgid", this.GetName(), msgid)
			return
		}

		//todo
		//goroutine
		//数据库查询客户端登陆信息函数
	}
}

//agent 进入场景
func (this *LoginSpace) LeaveSpace(ag *agent.Agent) {
	v, ok := this.Agents[ag.Aid]
	if ok {
		v.OnLeaveSpace()
		delete(this.Agents, v.Aid)
	}
}

//agent退出场景
func (this *LoginSpace) EnterSpace(ag *agent.Agent) {
	_, ok := this.Agents[ag.Aid]
	if ok {
		return
	}
	this.Agents[ag.Aid] = ag
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

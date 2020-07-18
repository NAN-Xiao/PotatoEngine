package space

import "potatoengine/src/agent"

type ISpace interface {

	GetName() string
	GetID() int32
	LeaveSpace(ag *agent.Agent)
	EnterSpace(ag *agent.Agent)
	//按时间间隔调用
	Tick()
	//不按时间同步调用
	Process()
}


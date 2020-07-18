package space

import "potatoengine/src/agent"

type ISpace interface {
	//todo 重新封装可能
	LeaveSpace(ag *agent.Agent)
	EnterSpace(ag *agent.Agent)


	//暂时不动
	SetGameID(gid int32)
	GetGameID()int32
	GetName() string
	GetID() int32
	//开始启动调用
	OnStart()
	//按时间间隔调用
	Tick()
	//不按时间同步调用
	Process()
}


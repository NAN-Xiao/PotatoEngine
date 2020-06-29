package space

import "potatoengine/src/agent"

type ISpace interface {
	Process()
	GetName() string
	GetID() int32
	LeaveSpace(ag *agent.Agent)
	EnterSpace(ag *agent.Agent)
}

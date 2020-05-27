package space

import "potatoengine/src/agent"

type ISpace interface {
	Process()
	GetName() string
	LeaveSpace(ag* agent.Agent)
	EnterSpace(ag *agent.Agent)
}

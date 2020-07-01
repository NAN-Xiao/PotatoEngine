package space

import (
	"potatoengine/src/agent"
	"potatoengine/src/netmessage"
)

type BaseSpace struct {
	SpaceID    int32
	Spacename  string
	Agents     map[uint32]*agent.Agent
	Spacechanl chan netmessage.ServerMsgPackage
}

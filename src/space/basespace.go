package space

import (
	"potatoengine/src/agent"
	"potatoengine/src/message"
)

type BaseSpace struct {
	SpaceID    uint32
	Spacename  string
	Agents     map[uint32]*agent.Agent
	Spacechanl chan *message.MsgPackage
}

//func (this *BaseSpace) Process() {
//
//}
//
//func (this *BaseSpace)GetName() string  {
//
//}
//func NewSpace(name string) *BaseSpace {
//
//}

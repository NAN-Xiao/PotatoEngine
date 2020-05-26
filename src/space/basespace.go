package space

import "potatoengine/src/message"

type BaseSpace struct {
	SpaceID    uint32
	Spacename  string
	Spacechanl chan message.Messsage
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

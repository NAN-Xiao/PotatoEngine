package space

import "potatoengine/src/message"

type BaseSpace struct {
	_name string
	_rch  map[int]chan *message.Messsage
	_wch  map[int]chan *message.Messsage
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

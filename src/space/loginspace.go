package space

import (
	"fmt"
	"potatoengine/src/message"
	"reflect"
)

type LoginSpace struct {
	BaseSpace
}

func (this *LoginSpace) Process() {

	fmt.Println("%s", reflect.TypeOf(this))
}

func (this *LoginSpace) GetName() string {
	return this._name
}

func NewLoginSpace(name string) ISpace {
	sp := &LoginSpace{
		BaseSpace{
			_name: name,
			_rch:  make(map[int]chan *message.Messsage),
			_wch:  make(map[int]chan *message.Messsage),
		},
	}
	return sp
}

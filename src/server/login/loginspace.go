package login

import "potatoengine/src/space"

type LoginSpace struct {
	_base space.BaseSpace
}

func (this *LoginSpace) Process() {

}
func (this *LoginSpace) GetName() string {
	//base :=&this._base
}

func NewLoginSpace(name string) *LoginSpace {
	sp := &LoginSpace{struct{ _name string }{_name: name}}
	return sp
}

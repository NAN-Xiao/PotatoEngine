package login

import "potatoengine/src/space"

type LoginSpace struct {
	_base space.BaseSpace
}

func (this *LoginSpace) Process() {

}
func (this *LoginSpace) GetName() string {
	//return this._base._name
}

func NewSpace(name string) *LoginSpace {
	sp := &LoginSpace{struct{ _name string }{_name: name}}
	return sp
}

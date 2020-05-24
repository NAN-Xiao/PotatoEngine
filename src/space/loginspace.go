package space

import (
	"fmt"
	"reflect"
)

type LoginSpace struct {
	BaseSpace
}

func (this *LoginSpace) Process() {
	fmt.Println("%s", reflect.TypeOf(this))
}

func (this *LoginSpace) GetName() string{
	return  this._name
}
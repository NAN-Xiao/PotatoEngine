package main

import (
	"fmt"
	"potatoengine/src/message"
	"potatoengine/src/space"
	"reflect"
)

type LoginSpace struct {
	space.BaseSpace
}

func (this *LoginSpace) Process() {
	go func() {
		for {
			value, ok := <-this.Spacechanl
			if ok == false {
				continue
			}
			if value.GetID() != message.Msg_Login {
				continue
			}
			//todo
			//goroutine
			//数据库查询客户端登陆信息函数
		}
	}()
	fmt.Println("%s", reflect.TypeOf(this))
}



func (this *LoginSpace) GetName() string {
	return this.Spacename
}

func NewLoginSpace(name string) space.ISpace {
	sp := &LoginSpace{struct {
		SpaceID    uint32
		Spacename  string
		Spacechanl chan message.Messsage
	}{SpaceID: 0, Spacename: "Login", Spacechanl: make(chan message.Messsage)}}
	return sp
}

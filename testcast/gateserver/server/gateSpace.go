package main

import (
	"potatoengine/src/netmessage"
	message "potatoengine/src/netmessage/pbmessage"
	"potatoengine/src/space"
	"reflect"
)

type GateSpace struct {
	space.BaseSpace
	space.ISpace
}

func (this *GateSpace) GetSpace() *space.BaseSpace {
	return &this.BaseSpace
}

//Ispace
func (this *GateSpace) OnStart() {
	println("gate space started")
	println(this.BaseSpace.Spacename)
}
func (this *GateSpace) Process() {
	for {
		if len(this.Entitys) < 0 {
			continue
		}
		for _, e := range this.Entitys {
			if e == nil {
				continue
			}
			entity, err := e.GetEntity()
			if err != nil {
				continue
			}
			msg := entity.Read()
			if msg != nil {
				m, err := netmessage.ExcutePBMessage(msg)
				if err != nil {
					continue
				}
				if reflect.TypeOf(m) == reflect.TypeOf(message.CheckTokenResult{}) {
					result, ok := m.(message.CheckTokenResult)
					if ok && result.Result == true {
						//todo 登录成功进入游戏
					}
				}
				entity.Write(m)
			}
		}
	}
}
func (this *GateSpace) Tick() {

	//println("gate space tick;entitys len ",len(this.Entitys))
}

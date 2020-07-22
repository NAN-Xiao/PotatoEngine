package main

import (
	"potatoengine/src/account"
	"potatoengine/src/netmessage"
	message "potatoengine/src/netmessage/pbmessage"
	"potatoengine/src/space"
)
//登录验证
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
			if e != nil {
				continue
			}
			ac, ok := e.(*account.Account)
			if ok == false {
				continue
			}
			entity := ac.GetEntity()
			if &entity == nil {
				continue
			}
			msg := entity.Read()
			if msg != nil {
				m, err := netmessage.ExcutePBMessage(msg)
				if err != nil {
					continue
				}
				if ac.Longin == false {
					result, ok := m.(message.CheckTokenResult)
					if ok && result.Result == true {
						//todo 登录成功进入游戏 从当前space移除account

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

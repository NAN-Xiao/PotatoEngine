package account

import (
	"potatoengine/src/client"
	"potatoengine/src/entity"
	"potatoengine/src/space"
)

type Account struct {
	Entity entity.Entity
	Longin bool
}



func (this *Account) GetEntityID() int32 {
	return this.Entity.GetEntityID()
}
func (this *Account) GetSpaceID() int32 {
	return this.GetSpaceID()
}
//进入场景
func (this *Account) EnterSpace(spaceID int32) {
	space:=space.GetSpace(spaceID)
	if space==nil{
		return
	}
}
//退出场景
func (this *Account) LeaveSpace(spaceID int32) {

}

func NewAccount(cl *client.Client) *Account {
	ac := &Account{
		Client:     cl,
		MsgChannel: make(chan interface{}, 2),
	}
	return ac
}
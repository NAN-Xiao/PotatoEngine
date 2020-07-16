package account

import (
	"fmt"
	"potatoengine/src/client"
)

type Account struct {
	Longin bool
	Client  *client.Client
}



func (this *Account) GetEntityID() int32 {
	return this._entityID
}
func (this *Account) GetSpaceID() int32 {
	return this._spaceID
}
//进入场景
func (this *Account) EnterSpace(spaceID int32) {

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
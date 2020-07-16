package account

import (
	"fmt"
	"potatoengine/src/client"
	"potatoengine/src/entity"
)

type Account struct {
	MsgChannel chan interface{}
	Longin bool

	Entity  entity.Entity
}
//从account读消息
func (this *Account) ReadMsg() (interface{}, error) {

	msg := <-this.MsgChannel
	if msg != nil {
		return msg, nil
	}
	return nil, fmt.Errorf("msg channel is nil")
}
//往account写消息
func (this *Account) WriteMsg(msg interface{}) {
	this.MsgChannel <- msg
}

func NewAccount(cl *client.Client) *Account {
	ac := &Account{
		Client:     cl,
		MsgChannel: make(chan interface{}, 2),
	}
	return ac
}

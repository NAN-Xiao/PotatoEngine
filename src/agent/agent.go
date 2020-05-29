package agent

import (
	"potatoengine/src/client"
	"potatoengine/src/message"
)

type Agent struct {
	//对应的client
	_client *client.Client
	//agent id
	_playerid uint32
	//当前角色所在场景的id
	_spaceID    uint32
	Sendchan    chan *message.MsgPackage
	Receivechan chan *message.MsgPackage
}

//得到clientid
func (this *Agent) GetUserID() uint32 {
	return this._client.UserID
}

//得到当前agnet的playerid
func (this *Agent) GetPlayerID() uint32 {
	return this._playerid
}

func (this *Agent) SendMessage(msgPackage *message.MsgPackage) {
	//todo
	//把当前消息打包成网络消息发送给client
	//this._client.Send(msgPackage)
}

//进入场景
func (this *Agent) OnEnterSpace() {

}

//退出场景
func (this *Agent) OnLeaveSpace() {

}

func NewAgent() *Agent {
	ag := &Agent{
		_client:     nil,
		_playerid:   0,
		Sendchan:    make(chan *message.MsgPackage, 20),
		Receivechan: make(chan *message.MsgPackage, 20),
	}
	return ag
}

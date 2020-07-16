package agent

import (
	"potatoengine/src/entity"
	"potatoengine/src/netmessage"
)

type Agent struct {
	_playerID   int32
	_spaceID    int32
	WriteChanel chan *netmessage.ServerMsgPackage
	ReadChanel  chan *netmessage.ServerMsgPackage

	Entity entity.Entity
}

//得到当前agnet的playerid
func (this *Agent) GetPlayerID() int32 {
	return this._playerID
}

func (this *Agent) WriteMessage(msgPackage *netmessage.ServerMsgPackage) {
	//todo
	//把当前消息打包成网络消息发送给client
	//this._client.Send(msgPackage)
}
func (this *Agent) ReadMessage(msgPackage *netmessage.ServerMsgPackage) {
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
		_playerID:   0,
		_spaceID:    0,
		WriteChanel: make(chan *netmessage.ServerMsgPackage, 20),
		ReadChanel:  make(chan *netmessage.ServerMsgPackage, 20),
	}
	return ag
}

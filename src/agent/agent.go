package agent

import (
	"potatoengine/src/client"
	"potatoengine/src/entity"
	"potatoengine/src/netmessage"
)

type Agent struct {
	_entityID   int32
	_spaceID    int32
	WriteChanel chan *netmessage.ServerMsgPackage
	ReadChanel  chan *netmessage.ServerMsgPackage
	Client      *client.Client
	entity.IEntity
}

//得到当前agnet的playerid
func (this *Agent) GetEntityID() int32 {
	return this._entityID
}
func (this *Agent) GetSpaceID() int32 {
	return this._spaceID
}
//进入场景
func (this *Agent) EnterSpace(spaceID int32) {

}

//退出场景
func (this *Agent) LeaveSpace(spaceID int32) {

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

func NewAgent(cl *client.Client) *Agent {
	ag := &Agent{
		_entityID:   0,
		_spaceID:    0,
		Client:      cl,
		WriteChanel: make(chan *netmessage.ServerMsgPackage, 20),
		ReadChanel:  make(chan *netmessage.ServerMsgPackage, 20),
	}
	return ag
}

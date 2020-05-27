package agent

import (
	"potatoengine/src/client"
	"potatoengine/src/message"
)

type Agent struct {
	//client id
	_cid uint32
	//agent id
	_aid        uint32
	Sendchan    chan *message.MsgPackage
	Receivechan chan *message.MsgPackage
}

//得到clientid
func (this *Agent) GetCLientID() uint32 {
	return this._cid
}

//得到Agentid
func (this *Agent) GetAgentID() uint32 {
	return this._aid
}

func (this *Agent) Run() {
	for {
		pkg := <-this.Sendchan
		if pkg == nil {
			continue
		}
		client.GetClientMgr().GetClient()
	}
}

//进入场景
func (this *Agent) OnEnterSpace() {

}

//退出场景
func (this *Agent) OnLeaveSpace() {

}

func NewAgent() *Agent {
	ag := &Agent{
		_cid:        0,
		_aid:        0,
		Sendchan:    make(chan *message.MsgPackage, 20),
		Receivechan: make(chan *message.MsgPackage, 20),
	}
	return ag
}

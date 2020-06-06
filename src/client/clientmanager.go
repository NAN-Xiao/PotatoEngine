package client

import (
	"fmt"
	"potatoengine/src/message"
)

//ClientMgr
type ClientMgr struct {
	_init       bool
	_clients    map[uint32]*Client
	_tempclient []*Client
}

var instance *ClientMgr

func (mgr *ClientMgr) Initialize() {
	//mgr._clients := list.New()
}

//根据id返回client
func (this *ClientMgr) GetClient(cid uint32) *Client {
	v, ok := this._clients[cid]
	if ok {
		return v
	}
	return nil
}

//当有连接。添加持有的客户端
func (this *ClientMgr) AddClient(cl *Client) {

	if this._tempclient == nil {
		return
	}
	for _, v := range this._tempclient {
		if v == cl {
			return
		}
	}
	this._tempclient = append(this._tempclient, cl)
	cl.OnConnection()
	fmt.Println("a client connected")
}

//删除持有的客户端
func (this *ClientMgr) RemoveCLient(cl *Client) {
	for _, v := range this._tempclient {
		if v == cl {
			//todo 从数组删除
			return
		}
	}
	if this._clients == nil {
		return
	}
	for k := range this._clients {
		if this._clients[k] == cl {
			delete(this._clients, k)
			return
		}
	}
}

func (this *ClientMgr) RemoveCLientByID(cid uint32) {
	v, o := this._clients[cid]
	if o == true {
		delete(this._clients, v.UserID)
	}
}

func GetClientMgr() *ClientMgr {
	if instance == nil || instance._init == false {
		instance = &ClientMgr{
			_init:       true,
			_clients:    make(map[uint32]*Client),
			_tempclient: make([]*Client, 0),
		}
	}
	return instance
}
//广播消息
func (mgr *ClientMgr) BroadcastMessage(msg *message.Messsage) {

	if mgr._clients == nil {
		return
	}
	for c := range mgr._clients {
		mgr._clients[c].Send(msg)
	}
}

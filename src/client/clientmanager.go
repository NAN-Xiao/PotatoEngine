package client

import (
	"container/list"
)

//ClientMgr
type ClientMgr struct {
	_init    bool
	_clients list.List
}

var inst *ClientMgr

func (mgr *ClientMgr) Initialize() {
	//mgr._clients := list.New()
}

//取出底部元素 先入先出
func (mgr *ClientMgr) Pop() *Client {
	if mgr._init == false {
		return nil
	}
	item := mgr._clients.Front()
	if item == nil {
		return nil
	}
	client := item.Value.(*Client)
	return client
}

//当有连接。添加持有的客户端
func (mgr *ClientMgr) AddClient(cl *Client) {

	for v := mgr._clients.Front(); v != nil; v = v.Next() {
		if v.Value == cl {
			return
		}
	}
	mgr._clients.PushBack(cl)
}

//删除持有的客户端
func (mgr *ClientMgr) RemoveCLient(cl *Client) {

	for v := mgr._clients.Front(); v != nil; v = v.Next() {
		if v.Value == cl {
			mgr._clients.Remove(v)
		}
	}
}
func GetClientMgr() *ClientMgr {
	if inst._init == false {
		inst = &ClientMgr{
			_init:    true,
			_clients: *list.New(),
		}
	}
	return inst
}

func (mgr *ClientMgr) BroadcastMessage() {

	for item := mgr._clients.Front(); nil != item; item = item.Next() {
		cl := Client(item.Value)
	}
}

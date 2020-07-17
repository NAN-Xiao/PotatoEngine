package client

import "potatoengine/src/logService"

var (
	_allClients  map[int32]*Client
	_clientIndex int32
)

func init() {
	_allClients = make(map[int32]*Client, 4096)
	_clientIndex=0
}
//全局持有客户端
func AddClient(cl *Client) {

	for _, client := range _allClients {
		if client == cl {
			logService.Log("the client is existing")
			return
		}
	}
	_clientIndex+=1
	cl.ConnID=_clientIndex
	_allClients[_clientIndex]=cl
}
func DeleteClient(cl *Client)  {
	if c,_:=_allClients[cl.ConnID];c!=nil{
		delete(_allClients,c.ConnID)
	}
	cl.Conn.Close()
}

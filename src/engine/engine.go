package engine

import "potatoengine/src/server"

var servers map[server.E_ServerNames]server.BaseServer

func AddServer(sr server.BaseServer) {
	if servers == nil {
		servers = make(map[server.E_ServerNames]server.BaseServer, 0)
	}
	_, ok := servers[sr.Name]
	if ok {
		return
	}
	servers[sr.Name] = sr
}

func RemoveServer(sr server.BaseServer) {
	if servers != nil {
		_, ok := servers[sr.Name]
		if ok {
			delete(servers, sr.Name)
			return
		}
	}
	println("remove ser fail")
}

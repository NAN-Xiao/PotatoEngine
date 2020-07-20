package server

var servers map[E_ServerNames] *BaseServer

func AddServer(sr *BaseServer) {
	if servers == nil {
		servers = make(map[E_ServerNames]*BaseServer, 0)
	}
	_, ok := servers[sr.Name]
	if ok {
		return
	}
	servers[sr.Name] = sr
}

func RemoveServer(sr *BaseServer) {
	if servers != nil {
		_, ok := servers[sr.Name]
		if ok {
			delete(servers, sr.Name)
			return
		}
	}
	println("remove ser fail")
}

func Start()  {
	if servers ==nil|| len(servers)<=0{
		println("engine not have any service")
	}
	for i:=range servers {
		servers[i].Run()
	}
}
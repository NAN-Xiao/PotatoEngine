package listener

import (
	"net"
	"potatoengine/src/account"
	"potatoengine/src/entity"
	connect2 "potatoengine/src/netWork/connect"
	"potatoengine/src/space"
)

type TcpListener struct {
	netWork string
	addr    string
	i       IListener
}

func (this *TcpListener) Listen() {
	go func() {
		addr, err := net.ResolveTCPAddr(this.netWork, this.addr)
		if err != nil {
			println(err)
			return
		}
		lisener, err := net.ListenTCP(this.netWork, addr)
		for {
			c, err := lisener.AcceptTCP()
			if err != nil {
				println(err)
				return
			}
			ac := new(account.Account)
			connect := connect2.NewTcpConnection(c)
			sp := space.GetSpaceByName("GateSpace")
			if sp == nil {
				c.Close()
				break
			}
			//当有客户端链接 建立一个account对象先放入到gate中
			ac.CreatEntity(connect)
			entity.RegistEntity(ac)
			ac.Connect()
			sp.GetSpace().EnterSpace(ac)
		}
	}()

}

func NewTcpListener(network string, add string) *TcpListener {
	listen := &TcpListener{
		netWork: network,
		addr:    add,
	}
	return listen
}

package netWork

import (
	"net"
	"potatoengine/src/account"
	"potatoengine/src/entity"
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
			connect:=NewTcpConnection(c)
			ac.CreatEntity(connect)
			entity.RegistEntity(ac)
			ac.Connect()
			sp:=space.GetSpaceByName("GateSpace")
			ac.EnterSpace(sp.GetSpace().SpaceID)
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

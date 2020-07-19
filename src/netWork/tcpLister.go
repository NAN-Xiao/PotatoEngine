package netWork

import (
	"net"
	"potatoengine/src/account"
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

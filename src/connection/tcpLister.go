package connection

import (
	"net"
	"potatoengine/src/client"
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
			cl := client.NewClient(c)
			cl.Conn_id.Set(client.GenConnID())
			client.AddClient(cl)
			cl.Connect()
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

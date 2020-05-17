package login

import (
	"fmt"
	"net"
	"potatoengine/src/client"
	"potatoengine/src/server"
)

type LoginServer struct {
	server.IServer
}

func (ls *LoginServer) Initialize() {

}
func (ls *LoginServer) Begin() {

	addr, err := net.ResolveTCPAddr("0.0.0.0", "8999")
	if err != nil {
		fmt.Println("loginserver start error from resolve addr")
		return
	}
	lisenter, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println("listener err")
		return
	}
	defer lisenter.Close()
	for {
		tcpConn, err := lisenter.AcceptTCP() //阻塞，当有客户端连接时，才会运行下面
		if err != nil {
			fmt.Println("tcpListener error :", err)
			continue
		}
		cl := client.NewClient(tcpConn)
		client.GetClientMgr().AddClient(cl)
	}

}

func (ls *LoginServer) Start() {
	ls.Begin()
}
func (ls *LoginServer) Stop() {

}

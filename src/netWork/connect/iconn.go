package connect

import "net"

type IConn interface {

	GetID() ConnID
	//网络收发
	Receive(chan interface{})
	Send(interface{})
	//关闭
	Close()
	IsClosed() bool
	GetRemoteAddr() net.Addr
}

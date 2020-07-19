package netWork

import "net"

type IConn interface {

	GetID() ConnID
	//网络收发
	Receive(chan interface{})
	Send()
	//内部收发
	Read() interface{}
	Write(interface{})
	//关闭
	Close()
	IsClosed() bool
	GetRemoteAddr() net.Addr

}

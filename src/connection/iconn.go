package connection

import "net"

type IConn interface {

	//网络收发
	Receive()
	Send()
	//内部收发
	Read() interface{}
	Write(interface{})
	//关闭
	Close()
	IsClosed() bool
	GetRemoteAddr() net.Addr

}

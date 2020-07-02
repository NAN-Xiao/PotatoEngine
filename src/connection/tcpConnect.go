package connection

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"potatoengine/src/netmessage"
)

type TcpConnect struct {
	conn *net.TCPConn
}

func (this *TcpConnect) Read()(l int,err error) {
	go func() {
		for {
			var buf = make([]byte, 4)
			_, err := io.ReadFull(this.conn,buf)
			if err != nil {
				fmt.Println(err)
				continue
			}
			len := binary.BigEndian.Uint32(buf)
			buf = make([]byte, len)
			_, err = io.ReadFull(this.conn,buf)
			if err!=nil{
				continue
			}
			id, obj := netmessage.UnPackNetMessage(buf)
			if id < 0 || obj == nil {
				//消息错误
				continue
			}
			//todo 接受数据分发消息
		}
	}()
	return
}
func (this *TcpConnect) Write(data []byte) {

}
func (this *TcpConnect) Close() bool {

}
func (this *TcpConnect) Listen() {

	addr, err := net.ResolveTCPAddr("tcp", "0,0,0,0:9000")
	if err != nil {
		return
	}
	lisener, err := net.ListenTCP("tcp", addr)
	c, err := lisener.AcceptTCP()
	this.conn = c
	//阻塞接受消息s
	this.Read()
}

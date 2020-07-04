package connection

import (
	"encoding/binary"
	"io"
	"net"
	"potatoengine/src/client"
	"potatoengine/src/netmessage"
	// "potatoengine/src/netmessage/pbmessage"
)

type TcpConnect struct {
	Clients []client.Client
}

func (this *TcpConnect) Read() (l int, err error) {
	go func() {

	}()
	return
}
func (this *TcpConnect) Write(data []byte) {

}
func (this *TcpConnect) Close() bool {
	return false
}

//整个server监听
func (this *TcpConnect) Listen() {
	go func() {
		addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:9000")
		if err != nil {
			println(err)
			return
		}
		lisener, err := net.ListenTCP("tcp", addr)
		for {
			c, err := lisener.AcceptTCP()
			if err != nil {
				println(err)
				return
			}
			println("new client")
			if this.Clients==nil{
				this.Clients=make([client.Client],0)
			}
			cl:=client.NewClient(c)
			this.Clients=append this.Clients
			go func(conn *net.TCPConn) {
				println("tcp listening")
				for {
					var buf = make([]byte, 4)
					n, err := io.ReadFull(conn, buf)
					if err == io.EOF {
						break
					}
					len := binary.BigEndian.Uint32(buf) - 4
					buf = make([]byte, len)
					n, err = io.ReadFull(conn, buf)
					if err == io.EOF || n < int(len) {
						break
					}
					id, object := netmessage.UnPackNetMessage(buf)
					if id < 0 || object == nil {
						//消息错误
						break
					}
					
					//接受数据分发消息 放到chanel缓冲
					// this.Chan <- object
				}
				conn.Close()
				println("CLose tcp con")
			}(c)
		}

	}()
}

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
	// Clients []*client.Client
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
			// println("new client")
			// if this.Clients == nil {
			// 	this.Clients = make([]*client.Client, 0)
			// }
			cl := client.NewClient(c)
			//this.Clients = append(this.Clients, cl)
			go func(conn *net.TCPConn, client *client.Client) {
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
					id, msg := netmessage.UnPackNetMessage(buf)
					if id < 0 || msg == nil {
						//消息错误
						break
					}
					cl.WriteToChanle(msg)
				}
				//移除持有的client 断开client的连接
				conn.Close()
				conn = nil
				cl = nil
				println("CLose tcp con")
			}(c, cl)
		}

	}()
}

package connection

import (
	"net"
)

type TcpConnect struct {
	conn []*net.TCPConn
}

func (this *TcpConnect) Read()(l int,err error) {
	go func() {

	}()
	return
}
func (this *TcpConnect) Write(data []byte) {

}
func (this *TcpConnect) Close() bool {
 return false
}
func (this *TcpConnect) Listen() {
	go func() {
		addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:9000")
		if err != nil {
			println(err)
			return
		}
		lisener, err := net.ListenTCP("tcp", addr)
		c, err := lisener.AcceptTCP()
		if err!=nil{
			return
		}
		println("new client")
		this.conn=append(this.conn,c)
		//go func(conn *net.TCPConn) {
		//	for {
		//		if this.conn==nil{
		//			continue
		//		}
		//		println("tcp listening")
		//		var buf = make([]byte, 4)
		//		_, err = io.ReadFull(conn,buf)
		//		if err != nil {
		//			fmt.Println(err)
		//			continue
		//		}
		//		len := binary.BigEndian.Uint32(buf)
		//		buf = make([]byte, len)
		//		_, err = io.ReadFull(conn,buf)
		//		if err!=nil{
		//			continue
		//		}
		//		id, obj := netmessage.UnPackNetMessage(buf)
		//		if id < 0 || obj == nil {
		//			//消息错误
		//			continue
		//		}
		//		//todo 接受数据分发消息
		//		println("message")
		//	}
		//}(c)
	}()
}

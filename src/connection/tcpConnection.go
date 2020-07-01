package connection

import "net"

type TcpConnect struct {
	conn *net.TCPConn
}

func (this *TcpConnect) Read(buf []byte) {
	this.conn.Read(buf)
}
func (this *TcpConnect) Write(data []byte) {

}
func (this *TcpConnect) Close() bool {

}
func (this *TcpConnect) Listen() {

}
func (this *TcpConnect)GetConnection() net.Conn {
	
}

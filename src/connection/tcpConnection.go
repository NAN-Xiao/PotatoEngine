package connection

import "net"

type TcpConnect struct {
	conn net.TCPConn
}

func (this *TcpConnect) Read() {
	go func() {
		for {
			var buf []byte
			this.conn.Read(buf)
			if len(buf) < 4 {
				continue
			}
		}
	}()

}
func (this *TcpConnect) Write(data []byte) {

}
func (this *TcpConnect) Close() bool {

}
func (this *TcpConnect) Listen() {

}

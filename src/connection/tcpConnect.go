package connection

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"potatoengine/src/netmessage"
)

type TcpConnect struct {
	Conn *net.TCPConn
	MsgChan chan interface{}
}
//接受消息放入队列
func (this *TcpConnect) Receive() error{
	for {
		var buf = make([]byte, 4)
		n, err := io.ReadFull(this.Conn, buf)
		if err == io.EOF {
			break
		}
		len := binary.BigEndian.Uint32(buf) - 4
		buf = make([]byte, len)
		n, err = io.ReadFull(this.Conn, buf)
		if err == io.EOF || n < int(len) {
			break
		}
		id, msg := netmessage.UnPackNetMessage(buf)
		if id < 0 || msg == nil {
			break
		}
		this.MsgChan<-msg
	}
	return fmt.Errorf("net msg process is error")
}
//从队列读取消息结构
func (this *TcpConnect)ReadMsg() interface{} {
	//if this.MsgChan==nil||len(this.MsgChan)<=0{
	//	return nil
	//}
	//return <-this.MsgChan
}
func (this *TcpConnect)GetMsgQue() chan interface{} {
	return  this.MsgChan
}
//发送消息
func  (this *TcpConnect)Write(data []byte) {

}
//关闭tcp链接
func (this *TcpConnect) Close() bool {
	return false
}


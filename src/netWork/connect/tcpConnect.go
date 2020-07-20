package connect

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"potatoengine/src/logService"
	"potatoengine/src/netmessage"
)

type TcpConnect struct {
	Conn     *net.TCPConn
	ConnID   ConnID
	SendChan chan interface{}
	iscloes  bool
}

func (this *TcpConnect) GetID() ConnID {
	return this.ConnID
}

//接受断言消息放入队列
func (this *TcpConnect) Receive(msgque chan interface{}) {
	for {
		if msgque == nil {
			return
		}
		if this.iscloes {
			break
		}
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
		msgque <- msg
	}
	logService.LogError(fmt.Sprintf("receive net msg  is error>>client id::%s", this.ConnID.Get()))
	this.Close()
}
func (this *TcpConnect) Send(msg interface{}) {
	for {
		if this.iscloes {
			break
		}
		msg := <-this.SendChan
		if msg == nil {
			continue
		}
		data, err := netmessage.PackageNetMessage(msg)
		if err != nil {
			continue
		}
		this.Conn.Write(data)
		continue
	}
}

//从队列读取消息结构
func (this *TcpConnect) Read() interface{} {
	if this.ReceiveChan == nil || len(this.ReceiveChan) <= 0 {
		return nil
	}
	return <-this.ReceiveChan
}

//本地调用缓存到发送消息队列
func (this *TcpConnect) Write(msg interface{}) {
	this.SendChan <- msg
}

//关闭tcp链接
func (this *TcpConnect) Close() {

	close(this.ReceiveChan)
	close(this.SendChan)
	this.Conn.Close()
	this.iscloes = true
}

//连接是否关闭
func (this *TcpConnect) IsClosed() bool {
	return this.iscloes
}
func (this *TcpConnect) GetRemoteAddr() net.Addr {
	return this.Conn.RemoteAddr()
}
func NewTcpConnection(con *net.TCPConn) *TcpConnect {
	tcp := &TcpConnect{
		Conn:        con,
		ConnID:      -1,
		ReceiveChan: make(chan interface{}, 128),
		SendChan:    make(chan interface{}, 128),
		iscloes:     false,
	}
	connectCount += 1
	tcp.ConnID.Set(connectCount)
	return tcp
}

var connectCount int32 = 0

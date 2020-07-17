package connection

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"potatoengine/src/netmessage"
)

type TcpConnect struct {
	Conn        *net.TCPConn
	ConnID      int32
	ReceiveChan chan interface{}
	SendChan    chan interface{}
	iscloes     bool
}

//接受断言消息放入队列
func (this *TcpConnect) Receive() error {
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
		this.ReceiveChan <- msg
	}
	return fmt.Errorf("net msg process is error")
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

//发送网络数据
func (this *TcpConnect) WriteToNet() {

	for {
		var msg interface{}
		msg = <-this.SendChan
		data, err := netmessage.PackageNetMessage(msg)
		if err != nil {
			return
		}
		this.Conn.Write(data)
	}
}
//关闭tcp链接
func (this *TcpConnect) Close() {

	close(this.ReceiveChan)
	close(this.SendChan)
	this.Conn.Close()
	this.iscloes=true
}
//连接是否关闭
func (this *TcpConnect)IsClosed()bool  {
	return this.iscloes
}

func NewTcpConnection(con *net.TCPConn, cid int32) *TcpConnect {
	tcp := &TcpConnect{
		Conn:        con,
		ConnID:      cid,
		ReceiveChan: make(chan interface{}, 128),
		SendChan:    make(chan interface{}, 128),
		iscloes:     false,
	}
	return tcp
}

package connection

import (
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"potatoengine/src/message"
)

type TcpConnnetion struct {
	_tcp_conn *net.TCPConn
	_msg_que  *message.MessageQue
	_closed   bool
	_len      uint32
	_wc       chan message.Messsage
	_rc       chan message.Messsage
}

//send消息外部接口。
////放到队列通过write发送客户端
//func (conn *TcpConnnetion) SendMessage(msg *message.Messsage) {
//
//	if conn._msg_que == nil || msg == nil {
//		fmt.Println("connection s msgque is nil")
//		return
//	}
//	conn._msg_que.PushBack(msg)
//}

func (conn *TcpConnnetion) Listen() {
	tempbuff := make([]byte, 0)
	nbuff := make([]byte, 2048)
	for {
		rlen, err := conn._tcp_conn.Read(nbuff)
		if err != nil {
			break
		}
		tempbuff := append(tempbuff, nbuff[0:rlen]...)
		if len(tempbuff) < 8 {
			continue
		}
		l := tempbuff[4:7]
		slen := binary.BigEndian.Uint32(l)
		if len(tempbuff) < int(slen)+8 {
			continue
		}

		//id := binary.BigEndian.Uint32(tempbuff[0:3])
		//data := tempbuff[8:slen]
		//msg := message.NewMessage(id, data)

		//重置tempbuff
		//dispatcher.DisposMessage(msg)
		tempbuff = make([]byte, 0)
	}

}

//从网络客户端接消息 处理粘包后塞到chanel
func (this *TcpConnnetion) ReadFormNet() bool {
	tempbuff := make([]byte, 0)
	buff := make([]byte, 2048)
	for {
		rlen, err := this._tcp_conn.Read(buff)
		if err != nil {
			break
		}
		tempbuff := append(tempbuff, buff[0:rlen]...)
		if len(tempbuff) < 8 {
			continue
		}
		l := tempbuff[4:7]
		slen := binary.BigEndian.Uint32(l)
		if len(tempbuff) < int(slen)+8 {
			continue
		}
		id := binary.BigEndian.Uint32(tempbuff[0:3])
		data := tempbuff[8:slen]
		msg := message.NewMessage(id, data)
		this._rc <- *msg
		tempbuff = make([]byte, 0)
	}
	fmt.Println("conent is break")
	return false
}

//读取通道并解析成消息 通过tcp向客户端发送
func (conn *TcpConnnetion) WriteToNet() {
	if conn._closed {
		fmt.Printf("client connect is closed")
		return
	}
	data := <-conn._wc
	pb, err := proto.Marshal((*data.GetData()))
	if err == nil {
		return
	}
	conn._tcp_conn.Write(pb)
}
//外部读取消息从chanel
func (conn *TcpConnnetion) ReadFromChannel() message.Messsage {
	msg := <-conn._rc
	return msg
}
//外部调用发送消息先放到chanel
func (conn *TcpConnnetion) WriteToChannel(msg message.Messsage) {
	conn._wc <- msg
}

//关闭连接
//tcp conn close()
//retun true
func (conn *TcpConnnetion) CloseConnection() bool {
	if conn._closed == false {
		conn._closed = true
	}
	conn._tcp_conn.Close()
	return conn._closed
}

//新建一个连/保持tcpconn
//新建的时候uid pid没有确认登陆前默认是0
func NewTcpConnection(tcpconn *net.TCPConn) *TcpConnnetion {
	con := &TcpConnnetion{
		_tcp_conn: tcpconn,
		_msg_que:  message.NewMessageQueue(10),
		//_buf:      make([]byte, 2048),
		_closed: false,
		_len:    4,
		_wc:     make(chan message.Messsage, 200),
		_rc:     make(chan message.Messsage, 200),
	}
	return con
}

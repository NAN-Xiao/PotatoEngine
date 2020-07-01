package connection

import (
	"encoding/binary"
	"net"
	"potatoengine/src/netmessage"
)

type Connnetion struct {
	_conn IConn
	_closed   bool
	_wc       chan netmessage.ServerMsgPackage
	_rc       chan netmessage.ServerMsgPackage
}

func (this *Connnetion) Listen() {
	tempbuff := make([]byte, 0)
	nbuff := make([]byte, 2048)
	addr,err:=net.ResolveTCPAddr("tcp","0,0,0,0:9000")
	if err!=nil{
		return
	}

	lisener,err:=net.ListenTCP("tcp",addr)
	c,err:=lisener.AcceptTCP()
	this._conn=&TcpConnect{
		conn:c,
	}
	var t_buf []byte
	var g_bug =make([]byte,0)
	for {
		this._conn.Read(t_buf)
		if len(t_buf)<=0{
			continue
		}
		if err != nil {
			break
		}
		tempbuff := append(g_bug, t_buf[0:len(t_buf)]...)
		if len(tempbuff) < 8 {
			continue
		}
		l := tempbuff[4:7]
		slen := binary.BigEndian.Uint32(l)
		if len(tempbuff) < int(slen)+8 {
			continue
		}

		tempbuff = make([]byte, 0)
	}

}

//从网络客户端接消息 处理粘包后塞到chanel
//func (this *Connnetion) Read() bool {
//	tempbuff := make([]byte, 0)
//	buff := make([]byte, 2048)
//	for {
//		rlen, err := this._conn.Read(buff)
//		if err != nil {
//			break
//		}
//		tempbuff := append(tempbuff, buff[0:rlen]...)
//		if len(tempbuff) < 8 {
//			continue
//		}
//		l := tempbuff[4:7]
//		slen := binary.BigEndian.Uint32(l)
//		if len(tempbuff) < int(slen)+8 {
//			continue
//		}
//		id := binary.BigEndian.Uint32(tempbuff[0:3])
//		data := tempbuff[8:slen]
//		msg := netmessage.PackMessagePackage(int32(id), 0, data)
//		this._rc <- *msg
//		tempbuff = make([]byte, 0)
//	}
//	fmt.Println("conent is break")
//	return false
//}

//读取通道并解析成消息 通过tcp向客户端发送
//func (conn *Connnetion) Write() {
//	if conn._closed {
//		fmt.Printf("client connect is closed")
//		return
//	}
//	data := <-conn._wc
//
//	msg, ok := data.GetMessage().(proto.Message)
//	if ok == false {
//		return
//	}
//
//	buf, err := proto.Marshal(msg)
//	if err != nil {
//		return
//	}
//	conn._tcp_conn.Write(buf)
//}

//外部读取消息从chanel
func (conn *Connnetion) ReadFromChannel() *netmessage.ServerMsgPackage {
	msg, ok := <-conn._rc
	if ok == true {
		return &msg
	}
	return nil
}

//外部调用发送消息先放到chanel
func (conn *Connnetion) WriteToChannel(msg netmessage.ServerMsgPackage) {
	conn._wc <- msg
}

//关闭连接
//tcp conn close()
//retun true
func (conn *Connnetion) Close() bool {
	if conn._closed == false {
		conn._closed = true
	}
	conn.Close()
	return conn._closed
}

//新建一个连/保持tcpconn
//新建的时候uid pid没有确认登陆前默认是0
func NewTcpConnection(t ConnType) *Connnetion {
	con := &Connnetion{
		//_msg_que:  netmessage.NewMessageQueue(10),
		//_buf:      make([]byte, 2048),
		_closed: false,
		_wc:     make(chan netmessage.ServerMsgPackage, 200),
		_rc:     make(chan netmessage.ServerMsgPackage, 200),
	}
	return con
}

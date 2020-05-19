package connection

import (
	"encoding/binary"
	"fmt"
	"net"
	"potatoengine/src/client"
)

type Connnetion struct {
	_tcp_conn net.TCPConn
	_buf      []byte
	_closed   bool
	_len      uint32
}

func (conn *Connnetion) Write(data []byte) {
	if conn._closed{
		fmt.Printf("client connect is closed")
		return
	}
	conn._tcp_conn.Write(data)
}

func (conn *Connnetion) Read() {
	for {
		len, err := conn._tcp_conn.Read(conn._buf)
		if uint32(len) < conn._len || err != nil {
			continue
		}
		stream := conn._buf[0:3]
		var head = binary.BigEndian.Uint32(stream)
		if uint32(len) < head {
			continue
		}
		stream := conn._buf[3 : head-1]
		//todo
		//解析登陆消息
		//账号
		//密码
		//查询数据库
		//返回登陆结果
		//push账号信息
		//push账号下角色信息
		//var message=ParsingLoginData(stream)
	}
}

//解析数据
func ParsingLoginData(data []byte) error {
	id := data[0:3]
	msgid := binary.BigEndian.Uint32(id)
	switch msgid {
	case 1:
		///登陆消息
		fmt.Println("login message")
	case 2:
		//其他消息
		fmt.Println("other message")
	default:
		return fmt.Errorf("message is not a loging msg")
		//错误消息
	}
	return nil
}

//关闭连接
//tcp conn close()
//retun true
func (conn *Connnetion) CloseConnection() bool {
	if conn._closed == false {
		conn._closed = true
	}
	conn._tcp_conn.Close()
	return conn._closed
}

func NewConnection(tcpconn *net.TCPConn) *Connnetion {
	con := &Connnetion{
		_tcp_conn: *tcpconn,
		_buf:      make([]byte, 2048),
		_closed:   false,
		_len:      4,
	}
	return con
}

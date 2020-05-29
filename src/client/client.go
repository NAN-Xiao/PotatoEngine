package client

import (
	"net"
	"potatoengine/src/connection"
	"potatoengine/src/dispatcher"
	"potatoengine/src/message"
)

type Client struct {
	UserID   uint32
	PlayerID uint32
	_conn    *connection.TcpConnnetion
}

func (cl *Client) Send(msg *message.Messsage) {
	if msg == nil {
		return
	}
	cl._conn.WriteToChannel(*msg)
}

///连接建立后开始循环读取写入派发消息
func (cl *Client) OnConnection() {
	go func() {
		for {
			cl._conn.ReadFormNet()
			cl._conn.WriteToNet()
			msg := cl._conn.ReadFromChannel()
			if &msg != nil {
				dispatcher.DispatcherMessage(cl.UserID, cl.PlayerID, &msg)
			}
		}
	}()
}

func NewClient(conn *net.TCPConn) *Client {
	_tconn := connection.NewTcpConnection(conn)
	if _tconn == nil {
		return nil
	}
	client := &Client{
		UserID:   0,
		PlayerID: 0,
		_conn:    _tconn,
	}
	return client
}

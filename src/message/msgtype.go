package message

type MsgType uint32

const (
	Msg_Login  MsgType = 1001
	Msg_Regist MsgType = 1002
	Msg_TEnter MsgType = 1003
)

var MsgHandle = map[int]string{
	1001: "Login",
	1002: "Regist",
	1003: "Enter",
}

package message


///为了将来容错。msgid从10001开始
type MsgID uint32
const (
	Msg_Login  MsgID = 1001
	Msg_Regist MsgID = 1002
	Msg_TEnter MsgID = 1003
)

//var MsgHandle = map[int]string{
//	1001: "Login",
//	1002: "Regist",
//	1003: "Enter",
//}

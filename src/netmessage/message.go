package netmessage

//用于服务器模块之间的通讯 和客户端用netmessagepackage
type ServerMsgPackage struct {
	//用户id
	Userid int32
	//角色id
	PlayerID int32
	//网络消息
	Msg interface{}
}

//打包成服务器间的消息模块
func PackMessagePackage(uid int32, pid int32, msg interface{}) *ServerMsgPackage {

	pack := &ServerMsgPackage{
		Userid:   uid,
		PlayerID: pid,
		Msg:      msg,
	}
	return pack
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

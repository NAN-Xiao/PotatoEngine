package netmessage

/////////////////////////
//用于服务器模块之间的通讯//
/////////////////////////
type MsgPackage struct {
	//用户id
	_uid uint32
	//角色id
	_playerID uint32
	//消息id
	_mid uint32
	//网络消息
	_msg interface{}
}

func (this *MsgPackage) GetMsgID() uint32 {
	return this._mid
}
func (this *MsgPackage)GetMessage()*interface{}  {
	return &this._msg
}

//打包成服务器间的消息模块
func PackMessagePackage(uid uint32, pid uint32, msg interface{}) *MsgPackage {
	pack := &MsgPackage{
		_uid:      uid,
		_playerID: pid,
		_msg:      msg,
	}
	return pack
}
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
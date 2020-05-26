package message

///服务器模块之间的通讯
type Package struct {
	_uid      uint32
	_playerID uint32
	_msg      Messsage
}

func NewPack(uid uint32, pid uint32, msg Messsage) *Package {
	pack := &Package{
		_uid:      uid,
		_playerID: pid,
		_msg:      msg,
	}
	return pack
}

//func Int32ToBytes(value uint32) []byte {
//
//	bytesBuffer := bytes.NewBuffer([]byte{})
//	binary.Write(bytesBuffer, binary.BigEndian, value)
//	return bytesBuffer.Bytes()
//}

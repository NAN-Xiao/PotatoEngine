package message

type messsage struct {
	_msgID uint32
	_len   uint32
	_body  []byte
}

func NewMessage(id uint32,data []byte) *messsage {
	msg := &messsage{
		_msgID: 0,
		_body:data,
	}
	msg._len=0
	return msg
}



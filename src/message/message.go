package message

type Messsage struct {
	_msgID uint32
	_len   uint32
	_body  []byte
}

func (msg *Messsage) GetData() *[]byte {
	return &msg._body
}

func NewMessage(id uint32, data []byte) *Messsage {
	msg := &Messsage{
		_msgID: 0,
		_body:  data,
	}
	msg._len = 0
	return msg
}

package message

type Messsage struct {
	_id   uint32
	_len  uint32
	_body []byte
}

func (msg *Messsage) GetData() *[]byte {
	return &msg._body
}

func (msg *Messsage) GetID() uint32 {
	return msg._id
}

func NewMessage(id uint32, data []byte) *Messsage {
	msg := &Messsage{
		_id:   id,
		_body: data,
	}
	msg._len = uint32(len(data))
	return msg
}

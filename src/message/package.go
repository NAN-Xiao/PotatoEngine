package message

import (
	"bytes"
	"encoding/binary"
)

type Package struct {
	_pk []byte
}

func NewPack(pk []byte) *Package {
	pack := &Package{_pk: pk}
	return pack
}
func Pack(msg Messsage) *Package {
	//data := msg.GetData()
	//len := Int32ToBytes(uint32(len(data) + 8))
	//id := Int32ToBytes(msgID)
	//pack := append(len, id[0:]...)
	//pack = append(pack, data[0:]...)
	//return NewPack(pack)
	return  nil
}

func UnPack(pk *Package) *Messsage {
	//todo 重新修改message这里要修改unpack
	//data := pk._pk
	//id := data[3:7]
	//body := data[8:]
	msg := &Messsage{
		//_len:0
		//_id: 0,
		//_pbmsg:
	}
	return msg
}

func Int32ToBytes(value uint32) []byte {

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, value)
	return bytesBuffer.Bytes()
}

package utility

import (
	"bytes"
	"encoding/gob"
)

func ConvertToBytes(data interface{}) []byte {
	buff:=new(bytes.Buffer)
	encoder:=gob.NewEncoder(buff)
	if err:=encoder.Encode(data);err!=nil{
		print(err)
		return nil
	}
	return buff.Bytes()
}

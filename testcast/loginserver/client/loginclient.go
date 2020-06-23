package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"net/http"
	msg "potatoengine/src/message/proto"
)


type LoginResponse struct {
	Msgid int
	Uid   int
	Token string
}


type UserInfo struct {
	Name string `json:"username"`
	Pass string `json:"password"`
}

func main() {

	var rp interface{}=&msg.LoginResponse{

	}
	des,ol:=rp.(descriptor.Message)
	if ol{
		_,md:=descriptor.MessageDescriptorProto(des)
		ext,_:=proto.GetExtension(md.GetOptions(), msg.E_ServerMsgID)
		s:=fmt.Sprint(ext)
		print(s)
	}


	url := "http://0.0.0.0:8999/login?a=1"
	// sc := []byte("client send")
	// buf := bytes.NewBuffer(sc)
	info := UserInfo{
		Name: "xiaonan",
		Pass: "123456",
	}
	//print(time.Now().Unix())
	data, _ := json.Marshal(info)
	//序列化json到二进制
	buf:=new(bytes.Buffer)
	binary.Write(buf,binary.LittleEndian,data)
	reqest, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return
	}
	defer reqest.Body.Close()
	var cc = http.Client{}
	if response,err:=cc.Do(reqest);err==nil{
		msg:=&LoginResponse{}
		decoder:=gob.NewDecoder(response.Body)
		decoder.Decode(msg)
		print("get msge \n")
		print(msg.Token)
	}
}

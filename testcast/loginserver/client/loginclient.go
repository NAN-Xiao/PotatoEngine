package main

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net/http"
	message "potatoengine/src/netmessage/pbmessage"
	//message "potatoengine/src/message/msg"
)

func main() {
	url := "http://0.0.0.0:8999/login?a=1"
	var rp =message.LoginResquest{
		Username: "xiaonan",
		Password: "123456",
	}
	data,_:=proto.Marshal(&rp)
	reqest, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	defer reqest.Body.Close()
	var cc = http.Client{}
	if response,err:=cc.Do(reqest);err==nil{
		msgresponse:=message.LoginResponse{}
		var data []byte
		response.Body.Read(data)
		proto.Unmarshal(data,&msgresponse)
		fmt.Println(msgresponse.Token)
	}
	fmt.Println("client end")
}

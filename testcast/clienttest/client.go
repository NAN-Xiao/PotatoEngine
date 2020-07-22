package main

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"potatoengine/src/netmessage"
	message "potatoengine/src/netmessage/pbmessage"
)

func main() {
	url := "http://0.0.0.0:8999/login?a=1"
	var rp = message.LoginResquest{
		Username: "xiaonan",
		Password: "123456",
	}
	data,_:= netmessage.PackageNetMessage(&rp)

	reqest, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	defer reqest.Body.Close()
	var cc = http.Client{}
	if response, err := cc.Do(reqest); err == nil {
		buf,_:=ioutil.ReadAll(response.Body)
		fmt.Println(buf)
		msgresponse := message.LoginResponse{}
		proto.Unmarshal(buf[8:], &msgresponse)
		fmt.Println("token:",msgresponse.Token)
	}
	fmt.Println("client end")
}

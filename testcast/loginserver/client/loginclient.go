package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"net/http"
)

type UserInfo struct {
	Name string `json:"username"`
	Pass string `json:"password"`
}

func main() {

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
	cc.Do(reqest)
}

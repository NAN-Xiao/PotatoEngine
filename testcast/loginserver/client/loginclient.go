package main

import (
	"bytes"
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

	data, _ := json.Marshal(info)
	reqest, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	defer reqest.Body.Close()
	var cc = http.Client{}
	cc.Do(reqest)
}

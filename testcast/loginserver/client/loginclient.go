package main

import (
	"bytes"
	"net/http"
)

func main() {

	url := "http://0.0.0.0:8999/login?a=1"
	sc := []byte("client send")
	buf := bytes.NewBuffer(sc)
	reqest, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return
	}
	defer reqest.Body.Close()
	var cc = http.Client{}
	cc.Do(reqest)
}

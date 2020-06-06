package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("begin")
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("addr is err")

		return
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Printf("conn is err:%v\n", err)
		return
	}
	go func() {
		for {

			conn.Write([]byte("hello"))
		}
	}()

	select {}
}

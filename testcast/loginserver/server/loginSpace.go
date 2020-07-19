package main

import (
	"fmt"
	_ "github.com/go-redis/redis"
	"net/http"
	"potatoengine/src/space"
)

//todo 定义消息类型（登陆请求 注册请求  账号密码错误 登陆错误 登陆成功,服务器异常）

//Json字段必须大写
type UserInfo struct {
	Id   int    `json:"id"`
	Name string `json:"username"`
	Pass string `json:"password"`
}

type LoginSpace struct {
	space.BaseSpace
	space.ISpace
}
func (this *LoginSpace) OnStart() {

}
//todo http监听返回登陆结果
func (this *LoginSpace) Process() {
	fmt.Println("Login Space start")
	http.HandleFunc("/login", LoginHandle)
	http.ListenAndServe("0.0.0.0:8999", nil)
}
func (this *LoginSpace)Tick()  {

}



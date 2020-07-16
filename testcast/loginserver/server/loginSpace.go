package main

import (
	"fmt"
	_ "github.com/go-redis/redis"
	"net/http"
	"potatoengine/src/agent"
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
}

//todo http监听返回登陆结果
func (this *LoginSpace) Process() {
	fmt.Println("Login Space start")
	http.HandleFunc("/login", LoginHandle)
	http.ListenAndServe("0.0.0.0:8999", nil)
}

//agent 进入场景
func (this *LoginSpace) LeaveSpace(ag *agent.Agent) {
	v, ok := this.Agents[ag.GetUserID()]
	if ok {
		v.OnLeaveSpace()
		delete(this.Agents, v.GetUserID())
	}
}
func (this *LoginSpace) GetID() int32 {
	return this.SpaceID
}

//agent退出场景
func (this *LoginSpace) EnterSpace(ag *agent.Agent) {
	_, ok := this.Agents[ag.GetUserID()]
	if ok {
		return
	}
	this.Agents[ag.GetUserID()] = ag
	ag.OnLeaveSpace()
}

func (this *LoginSpace) GetName() string {
	return this.Spacename
}
func (this *LoginSpace)Tick()  {

}
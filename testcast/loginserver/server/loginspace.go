package main

import (
	"crypto/md5"
	"encoding/binary"
	"potatoengine/src/netmessage"

	"fmt"
	_ "github.com/go-redis/redis"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"potatoengine/src/agent"
	message "potatoengine/src/netmessage/pbmessage"
	"potatoengine/src/space"
	"time"
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

func Login(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("login space is runing")
	var buf, err = ioutil.ReadAll(request.Body)
	fmt.Println(buf)
	if err != nil || buf == nil || len(buf) <= 0 {
		fmt.Println(err)
		return
	}
	//pb消息

	responsdata := buf[4:]
	var loginRequest = message.LoginResquest{}
	err = proto.Unmarshal(responsdata, &loginRequest)
	if err != nil {
		fmt.Println(err)
		return
	}
	d := buf[4:8]
	id := binary.LittleEndian.Uint32(d)
	fn := netmessage.GetProcessFuction(int32(id))
	if fn == nil {
		return
	}
	loginResponse, netError := fn(loginRequest)
	if netError != nil {
		data, _ := netmessage.UnCodePBNetMessage(&netError)
		if data == nil {
			data, _ = netmessage.DefaultNetErrorData()
		}
		writer.Write(data)
		return
	}
	msgdata, _ := netmessage.UnCodePBNetMessage(loginResponse)
	writer.Write(msgdata)
}

//todo http监听返回登陆结果
func (this *LoginSpace) Process() {
	fmt.Println("Login Space start")
	http.HandleFunc("/login", Login)
	http.ListenAndServe("0.0.0.0:8999", nil)
}

//根据时间 id password计算toke md5
func CenerateToken(name string, password string) (string, error) {
	if &name == nil || name == "" {
		err := fmt.Errorf("username is nil")
		return "", err
	}
	if password == "" || &password == nil {
		err := fmt.Errorf("password is nil")
		return "", err

	}
	st := fmt.Sprintf("%s%s%s", name, password, time.Now().Unix())
	data := []byte(st)
	mdstr := md5.Sum(data)
	return fmt.Sprintf("%x", mdstr), nil
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

//func NewLoginSpace(name string) space.ISpace {
//	sp := &LoginSpace{space.BaseSpace{
//		SpaceID:    0,
//		Spacename:  name,
//		Agents:     make(map[uint32]*agent.Agent),
//		Spacechanl: make(chan *netmessage.MsgPackage, 100),
//	}}
//	return sp
//}

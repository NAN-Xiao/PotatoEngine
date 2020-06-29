package main

import (
	"crypto/md5"

	"fmt"
	_ "github.com/go-redis/redis"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"potatoengine/src/agent"
	"potatoengine/src/db"
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
	var loginRequest=message.LoginResquest{}
	err=proto.Unmarshal(responsdata,&loginRequest)
	if err != nil {
		fmt.Println(err)
		return
	}
	//查询数据库
	var id int32
	var name string
	var pass string
	sql := db.GetSQLManager().GetSQL()
	if sql == nil || sql.Ping() != nil {
		//发送服务器错误消息
		var errinfo interface{} = &message.NetError{ErrorCode: message.EMsg_Error_DBClosed}
		data, _ := errinfo.(proto.Message)
		pb, _ := proto.Marshal(data)
		writer.Write(pb)
		return
	}

	myquery := fmt.Sprintf("SELECT * FROM `game`.`userinfo` WHERE `username` = '%s'AND`password`='%s'", loginRequest.Username, loginRequest.Password)
	rows := sql.QueryRow(myquery)
	//mysql错误
	if rows != nil {
		rows.Scan(&id, &name, &pass)
		if id <= 0 {
			var errinfo interface{} = &message.NetError{ErrorCode: message.EMsg_Error_UserInfo}
			data, _ := errinfo.(proto.Message)
			pb, _ := proto.Marshal(data)
			writer.Write(pb)
			return
		}
	}
	//生成token返回token
	//存到redis key是id value是token
	redis, erro := db.GetRedisManager().GetDB()
	//检查redis
	if erro == nil {
		var errinfo interface{} = &message.NetError{ErrorCode: message.EMsg_Error_DBClosed}
		data, _ := errinfo.(proto.Message)
		pb, _ := proto.Marshal(data)
		writer.Write(pb)
		return
	}
	//生成token
	var token string

	if token, err = CenerateToken(name, pass); err == nil {
		var errinfo interface{} = &message.NetError{ErrorCode: message.EMsg_Error_DBClosed}
		data, _ := errinfo.(proto.Message)
		pb, _ := proto.Marshal(data)
		writer.Write(pb)
		return
	}
	//把token写入redis
	if err = redis.Set(fmt.Sprintf("userid_%d", id), token, 0).Err(); err != nil {
		var errinfo interface{} = &message.NetError{ErrorCode: message.EMsg_Error_DBClosed}
		data, _ := errinfo.(proto.Message)
		pb, _ := proto.Marshal(data)
		writer.Write(pb)
		return
	}
	//所有验证通过 填充responsmsg内容 返回客户端登陆结果
	loginResponse:=message.LoginResponse{}
	loginResponse.Userid=id
	loginResponse.Token=token
	data,_:=proto.Marshal(&loginResponse)
	writer.Write(data)
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

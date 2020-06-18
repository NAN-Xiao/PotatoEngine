package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-redis/redis"
	"io/ioutil"
	"net/http"
	"potatoengine/src/agent"
	"potatoengine/src/db"
	"potatoengine/src/message"
	"potatoengine/src/space"
	"time"
)

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
	// print(request)
	buf, err := ioutil.ReadAll(request.Body)
	fmt.Println("Server response")
	print(buf)
	//panic(err)
	if err != nil || buf == nil || len(buf) <= 0 {
		print(err)
		return
	}
	//反序列化json
	info := UserInfo{}
	err = json.Unmarshal(buf, &info)
	if err != nil {
		print("Unmarshal err")
		return
	}
	sql := db.GetSQLManager().GetSQL()

	if sql == nil || sql.Ping() != nil {
		print("sql not started")
		return
	}
	myquery := fmt.Sprintf("SELECT * FROM `game`.`userinfo` WHERE `username` = '%s'AND`password`='%s'", info.Name, info.Pass) //, info.Pass)
	rows := sql.QueryRow(myquery)

	if rows == nil {
		print("not find \n")
		return
	}
	var id int
	var name string
	var pass string

	rows.Scan(&id, &name, &pass)
	if id < 0 {
		//todo
		//没有id或者id错误
	}
	//todo

	//生成token返回token
	//存到redis key是id value是token
	redis, erro := db.GetRedisManager().GetDB()
	if erro != nil {
		return
	}
	if _, err := redis.Ping().Result(); err != nil {
		print("redis server not connected")
		return
	}
	//返回登陆成功消息

	// if matching == true {
	// 	writer.Write([]byte{1})
	// 	return
	// }
	// writer.Write([]byte{0})
}

//todo http监听返回登陆结果
func (this *LoginSpace) Process() {
	http.HandleFunc("/login", Login)
	http.ListenAndServe("0.0.0.0:8999", nil)
	//fmt.Println("loging server started")
}

//根据时间 id password计算toke
func CenerateToken(name string, password string) (string, error) {
	if &name == nil || name == "" {
		err := fmt.Errorf("username is nil")
		return "", err
	}
	if password == "" || &password == nil {
		err := fmt.Errorf("password is nil")
		return "", err

	}
	token := fmt.Sprintf("%s%s%s", name, password, time.Now().Unix())
	print(token)
	return token, nil
}

//agent 进入场景
func (this *LoginSpace) LeaveSpace(ag *agent.Agent) {
	v, ok := this.Agents[ag.GetUserID()]
	if ok {
		v.OnLeaveSpace()
		delete(this.Agents, v.GetUserID())
	}
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

func NewLoginSpace(name string) space.ISpace {
	sp := &LoginSpace{space.BaseSpace{
		SpaceID:    0,
		Spacename:  name,
		Agents:     make(map[uint32]*agent.Agent),
		Spacechanl: make(chan *message.MsgPackage, 100),
	}}
	return sp
}

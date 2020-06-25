package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	_ "github.com/go-redis/redis"
	"io/ioutil"
	"net/http"
	"potatoengine/src/agent"
	"potatoengine/src/db"
	"potatoengine/src/message"
	"potatoengine/src/space"
	"potatoengine/src/utility"
	"time"
)

//todo 定义消息类型（登陆请求 注册请求  账号密码错误 登陆错误 登陆成功,服务器异常）

//登录请求
type LoginRequest struct {
}

//返回消息
type LoginResponse struct {
	Msgid int
	Uid   int
	Token string
}

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
	responsmsg := &LoginResponse{
		Msgid: 0,
		Uid:   0,
		Token: "",
	}
	var responsdata []byte
	buf, err := ioutil.ReadAll(request.Body)
	fmt.Println("Server response")
	print(buf)
	if err != nil || buf == nil || len(buf) <= 0 {
		responsmsg.Msgid = 0
		responsdata = utility.ConvertToBytes(responsmsg)
		writer.Write(responsdata)
		return
	}
	//todo 以后修改成二进制
	//反序列化json
	info := UserInfo{}
	if err = json.Unmarshal(buf, &info); err != nil {
		responsmsg.Msgid = 0
		responsdata = utility.ConvertToBytes(responsmsg)
		writer.Write(responsdata)
		return
	}
	//建立mysql连接 查询账号数据
	sql := db.GetSQLManager().GetSQL()
	if sql == nil || sql.Ping() != nil {
		print("sql not started")
		responsmsg.Msgid = 0
		responsdata = utility.ConvertToBytes(responsmsg)
		writer.Write(responsdata)
		return
	}
	var id int
	var name string
	var pass string
	myquery := fmt.Sprintf("SELECT * FROM `game`.`userinfo` WHERE `username` = '%s'AND`password`='%s'", info.Name, info.Pass)
	rows := sql.QueryRow(myquery)
	if rows == nil {
		responsmsg.Msgid = 0
		responsdata = utility.ConvertToBytes(responsmsg)
		writer.Write(responsdata)
		print("not find \n")
		return
	}
	rows.Scan(&id, &name, &pass)
	if id < 0 {
		responsmsg.Msgid = 0
		responsdata = utility.ConvertToBytes(responsmsg)
		writer.Write(responsdata)
	}
	//生成token返回token
	//存到redis key是id value是token
	redis, erro := db.GetRedisManager().GetDB()
	//检查redis
	if erro == nil {
		if _, err := redis.Ping().Result(); err != nil {
			print("redis server not connected")
			responsmsg.Msgid = 0
			responsdata = utility.ConvertToBytes(responsmsg)
			writer.Write(responsdata)
			return
		}
	} else {
		responsmsg.Msgid = 0
		responsdata = utility.ConvertToBytes(responsmsg)
		writer.Write(responsdata)
		return
	}

	//生成token
	var token string
	if token, err = CenerateToken(name, pass); err != nil {
		responsmsg.Msgid = 0
		responsdata = utility.ConvertToBytes(responsmsg)
		writer.Write(responsdata)
		return
	}
	//把token写入redis
	if err = redis.Set(fmt.Sprintf("userid_%d", id), token, 0).Err(); err != nil {
		responsmsg.Msgid = 0
		responsdata = utility.ConvertToBytes(responsmsg)
		writer.Write(responsdata)
		return
	}
	//所有验证通过 填充responsmsg内容 返回客户端登陆结果
	responsmsg.Msgid = 0
	responsmsg.Token = token
	responsmsg.Uid = id
	responsdata = utility.ConvertToBytes(responsmsg)
	writer.Write(responsdata)
	//print("done")
}

//todo http监听返回登陆结果
func (this *LoginSpace) Process() {
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
		Spacechanl: make(chan *netmessage.MsgPackage, 100),
	}}
	return sp
}
